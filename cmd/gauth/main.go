package main

import (
	"encoding/base64"
	"fmt"
	"github.com/negronjl/libGoogle/pkg/gauth"
	"log"
	"os"
)

func main() {
	cjson := os.Getenv("CREDENTIALS_JSON")
	if cjson == "" {
		log.Fatal("Unable to load CREDENTIALS_JSON environment variable ")
	}
	scopes := os.Getenv("SCOPES")
	if scopes == "" {
		log.Fatal("Unable to load SCOPES environment variable ")
	}
	token, err := gauth.GetToken(cjson, scopes)
	if err != nil {
		log.Fatalf("Unable to get token: %v", err)
	}
	fmt.Printf("Token: [%v]\n", base64.StdEncoding.EncodeToString(token))
}
