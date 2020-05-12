package main

import (
	"fmt"
	"github.com/negronjl/libGoogle/pkg/gauth"
	"github.com/negronjl/libGoogle/pkg/gsuite"
	"log"
	"os"
)

func main(){
	userEmail := os.GetEnv("USER_EMAIL")
	if userEmail == "" {
		log.Fatal("Unable to load USER_EMAIL environment variable ")
	}
	cjson := os.Getenv("CREDENTIALS_JSON")
	if cjson == "" {
		log.Fatal("Unable to load CREDENTIALS_JSON environment variable ")
	}
	scopes := os.Getenv("SCOPES")
	if scopes == "" {
		log.Fatal("Unable to load SCOPES environment variable ")
	}
	authToken := os.Getenv("AUTH_TOKEN")
	if authToken == "" {
		log.Fatal("Unable to load AUTH_TOKEN environment variable ")
	}
	client := gauth.GetClient(cjson, scopes, authToken)
	if client == nil {
		log.Fatalln("Unable to instantiate http.Client")
	}
	backupCodes, err := gsuite.GetBackupCodes(client, userEmail)
	if err != nil {
		log.Fatalf("Error getting backup codes for %s: %v", userEmail, err)
	}
	fmt.Printf("Backup Codes for %s: %v", userEmail, backupCodes)
}
