package request

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
)

type Values map[string]interface{}

func (vals Values) String() string {
	b, _ := json.Marshal(vals)

	return string(b)
}

func (vals Values) Set(key string, value interface{}) {
	vals[key] = value
}

func (vals Values) Get(key string) interface{} {
	return vals[key]
}

func (vals Values) Delete(key string) {
	delete(vals, key)
}

func (vals Values) ToQueryParams() (map[string]string, error) {
	params := map[string]string{}
	for k, v := range vals {
		switch val := v.(type) {
		case string:
			params[k] = val
		case int:
			params[k] = strconv.Itoa(val)
		case int64:
			params[k] = strconv.FormatInt(val, 10)
		case float32:
			params[k] = strconv.FormatFloat(float64(val), 'f', -1, 32)
		case float64:
			params[k] = strconv.FormatFloat(val, 'f', -1, 64)
		case bool:
			params[k] = strconv.FormatBool(val)
		default:
			return map[string]string{}, errors.New(reflect.TypeOf(v).String() + " type can not be converted to string")
		}
	}

	return params, nil
}

func getPointer(v interface{}) interface{} {
	vv := reflect.ValueOf(v)
	if vv.Kind() == reflect.Ptr {
		return v
	}
	return reflect.New(vv.Type()).Interface()
}

func unmarshal(resType, val interface{}) (interface{}, error) {
	res := getPointer(resType)

	data, err := json.Marshal(val)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (vals Values) GetResult(resType interface{}) (interface{}, error) {
	return unmarshal(resType, vals)
}
