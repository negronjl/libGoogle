package main

import (
	"fmt"
	"github.com/negronjl/libGoogle/pkg/gauth"
	"github.com/negronjl/libGoogle/pkg/gsuite"
	"log"
	"os"
)

func main(){
	queryParam := "givenName:* familyName:*"
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
	users, err := gsuite.GetGsuiteUsers(client, queryParam)
	if err != nil {
		log.Fatalf("Unable to get users: %v", err)
	}
	for _, user := range users {
		accountStatus := "Active"
		if user.Suspended || user.Archived{
			accountStatus = "InActive"
		}
		fmt.Printf("User: %s %s (%s)  Account Status: %s\n", user.Name.GivenName, user.Name.FamilyName, user.PrimaryEmail, accountStatus)
	}
	fmt.Printf("Users downloaded: %v", len(users))

}
