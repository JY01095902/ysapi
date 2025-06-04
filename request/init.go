package request

import (
	"log"

	"gitlab.libratone.com/internet/ysapi.git/env"
)

func init() {
	log.Println("getting url root...")
	tenantId := env.MustGetEnv("YS_TENANT_ID")

	res, err := getURLRoot(tenantId)
	if err != nil {
		log.Printf("get url root failed: %v", err)
	}

	log.Printf("gateway url: %v", res.Data.GatewayURL)
	log.Printf("token url: %v", res.Data.TokenURL)

	tokenURLRoot = res.Data.TokenURL
	URLRoot = res.Data.GatewayURL

	refreshURLs()
}
