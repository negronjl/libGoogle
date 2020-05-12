package main

import (
	"fmt"
	"github.com/negronjl/libGoogle/pkg/gauth"
	"github.com/negronjl/libGoogle/pkg/gsuite"
	"log"
	"os"
)

func main() {
	queryParam := "name:*"
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
	groups, err := gsuite.GetGsuiteGroups(client, queryParam)
	if err != nil {
		log.Fatalf("Unable to get groups information: %v", err)
	}
	for _, group := range groups {
		fmt.Printf("Group: %s (%s)", group.Name, group.Email)
	}
}
