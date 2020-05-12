package main

import (
	"fmt"
	"github.com/negronjl/libGoogle/pkg/gauth"
	"github.com/negronjl/libGoogle/pkg/gsuite"
	"log"
	"os"
)

func main() {
	userEmail := os.Getenv("USER_EMAIL")
	if userEmail == "" {
		log.Fatalln("Unable to load the USER_EMAIL environment variable")
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
	userInfo, err := gsuite.GetGsuiteUser(client, userEmail)
	if err != nil {
		log.Fatalf("Unable to load user [%s] information: %v", err)
	}
	fmt.Printf("%s %s (%s)\n", userInfo.Name.GivenName, userInfo.Name.FamilyName, userInfo.PrimaryEmail)
}
