package ysapi

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func checkMessage(message EventMessage, secret string) bool {
	getSHA1 := func(secret, timestamp, nonce, ciphertext string) string {
		params := []string{secret, timestamp, nonce, ciphertext}
		sort.Strings(params)
		str := strings.Join(params, "")

		hexStr := ""
		digest := sha1.Sum([]byte(str))
		for i := 0; i < len(digest); i++ {
			shaHex := fmt.Sprintf("%x", digest[i]&0xFF)
			if len(shaHex) < 2 {
				hexStr += "0"
			}

			hexStr += shaHex
		}

		return hexStr
	}

	signature := getSHA1(secret, strconv.FormatInt(message.Timestamp, 10), message.Nonce, message.Ciphertext)

	return signature == message.Signature
}

type EventMessage struct {
	Signature  string `json:"msgSignature"`
	Timestamp  int64  `json:"timestamp"`
	Nonce      string `json:"nonce"`
	Ciphertext string `json:"encrypt"`
}

func (e EventMessage) String() string {
	b, _ := json.Marshal(e)

	return string(b)
}

func (e EventMessage) Decrypt(appKey, appSecret string) (EventPlaintext, error) {
	if !checkMessage(e, appSecret) {
		return EventPlaintext{}, errors.New("signature is error")
	}

	plaintext, suiteKey, err := decryptMessage(e.Ciphertext, appSecret)
	if err != nil {
		return EventPlaintext{}, errors.New("decrypt failed, error: " + err.Error())
	}

	if suiteKey != appKey {
		return EventPlaintext{}, errors.New("message is invalid")
	}

	var eventPlaintext EventPlaintext

	d := json.NewDecoder(bytes.NewReader([]byte(plaintext)))
	d.UseNumber()
	if err := d.Decode(&eventPlaintext); err != nil {
		return eventPlaintext, err
	}

	// if err := json.Unmarshal([]byte(plaintext), &eventPlaintext); err != nil {
	// 	return EventPlaintext{}, err
	// }

	return eventPlaintext, nil
}

type EventPlaintext struct {
	EventType     string   `json:"type"`
	EventId       string   `json:"eventId"`
	Timestamp     int64    `json:"timestamp"`
	TenantId      string   `json:"tenantId"`
	StaffIds      []string `json:"staffId"`
	DepartmentIds []string `json:"deptId"`
	UserIds       []string `json:"userId"`
	Content       string   `json:"content"`
}

func (e EventPlaintext) String() string {
	b, _ := json.Marshal(e)

	return string(b)
}

func (e EventPlaintext) UnmarshalContent() (EventContent, error) {
	var content EventContent

	d := json.NewDecoder(bytes.NewReader([]byte(e.Content)))
	d.UseNumber()
	if err := d.Decode(&content); err != nil {
		return content, err
	}

	// if err := json.Unmarshal([]byte(e.Content), &content); err != nil {
	// 	return EventContent{}, err
	// }

	return content, nil
}

type EventContent struct {
	AccessToken    string      `json:"access_token"`
	YHTUser        string      `json:"yhtUser"`
	TenantId       string      `json:"tenantId"`
	Archive        string      `json:"archive"`
	YHTAccessToken string      `json:"yht_access_token"`
	Fullname       string      `json:"fullname"`
	Id             json.Number `json:"id"`
	EventParams    struct {
		Data []Values `json:"data"`
	} `json:"eventParams"` // 库存分配事件的数据在这里
}

func (e EventContent) String() string {
	b, _ := json.Marshal(e)

	return string(b)
}

func (e EventContent) UnmarshalArchive() (Values, error) {
	if e.Archive == "" {
		return Values{}, nil
	}

	var val Values
	d := json.NewDecoder(bytes.NewReader([]byte(e.Archive)))
	d.UseNumber()
	if err := d.Decode(&val); err != nil {
		return val, err
	}

	// if err := json.Unmarshal([]byte(e.Archive), &val); err != nil {
	// 	return nil, err
	// }

	return val, nil
}

func makeAESKeyFromSecret(secret string) string {
	key := strings.ReplaceAll(secret, "-", "")
	if len(key) == 43 {
		return key
	}

	if len(key) > 43 {
		return key[0:43]
	}

	for len(key) < 43 {
		key += "0"
	}

	return key
}

func recoverNetworkBytesOrder(orderBytes []byte) int {
	sourceNumber := 0
	for i := 0; i < 4; i++ {
		sourceNumber <<= 8
		sourceNumber |= int(orderBytes[i]) & 0xFF
	}
	return sourceNumber
}

func pkcs7Decode(decrypted []byte) []byte {
	pad := int(decrypted[len(decrypted)-1])
	if pad < 1 || pad > 32 {
		pad = 0
	}

	return decrypted[:len(decrypted)-pad]
}

func decryptMessage(ciphertext, secret string) (string, string, error) {
	encodingAesKey := makeAESKeyFromSecret(secret)
	aesKey, err := base64.StdEncoding.DecodeString(encodingAesKey + "=")
	if err != nil {
		return "", "", err
	}

	cipherbyte, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", "", err
	}

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return "", "", err
	}

	iv := aesKey[0:16]
	mode := cipher.NewCBCDecrypter(block, iv)

	plainbyte := make([]byte, len(cipherbyte))
	mode.CryptBlocks(plainbyte, cipherbyte)

	decoding := pkcs7Decode(plainbyte)
	networkOrder := decoding[16:20]
	xmlLength := recoverNetworkBytesOrder(networkOrder)
	message := decoding[20 : 20+xmlLength]
	suiteKey := decoding[20+xmlLength:]

	return string(message), string(suiteKey), nil
}

func ParseEventMessage(message string) (EventMessage, error) {
	var eventMessage EventMessage
	d := json.NewDecoder(bytes.NewReader([]byte(message)))
	d.UseNumber()
	if err := d.Decode(&eventMessage); err != nil {
		return eventMessage, err
	}

	// err := json.Unmarshal([]byte(message), &eventMessage)

	return eventMessage, nil
}
