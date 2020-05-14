package main

import (
	"fmt"
	"github.com/negronjl/libGoogle/pkg/gauth"
	gmail2 "github.com/negronjl/libGoogle/pkg/gmail"
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

	authToken := os.Getenv("AUTH_TOKEN")
	if authToken == "" {
		log.Fatal("Unable to load AUTH_TOKEN environment variable ")
	}

	client := gauth.GetClient(cjson, scopes, authToken)
	if client == nil {
		log.Fatal("Unable to instantiate http.Client")
	}

	userEmail := os.Getenv("USER_EMAIL")
	if userEmail == "" {
		log.Fatal("Unable to load USER_EMAIL environment variable ")
	}

	labels, err := gmail2.ListLabels(client, userEmail)
	if err != nil {
		log.Fatalf("Error getting the list of labels from Gmail: %v", err)
	}
	for _, label := range labels {
		fmt.Printf("- %s\n", label)
	}

	query := os.Getenv("GMAIL_QUERY")
	if query == "" {
		log.Fatal("Unable to load GMAIL_QUERY environment variable ")
	}

	messages , err := gmail2.ListMessagesMatchingQuery(client, userEmail, query)
	if err != nil {
		log.Fatalf("Error getting messages from Gmail: %v", err)
	}

	for _, message := range messages {
		msgContent, err := gmail2.GetMessageById(client, userEmail, message.Id)
		if err != nil {
			fmt.Printf("Unable to get message id: %s\n", message.Id)
		} else {
			fmt.Printf("- %v\n", msgContent.Snippet)
		}

	}
	fmt.Printf("Label count: %d\n", len(labels))
	fmt.Printf("Message count: %d\n", len(messages))
	lastMessageId := messages[len(messages)-1].Id
	fmt.Printf("Last Message ID: %v\n", lastMessageId)
	lastMessage, err := gmail2.GetMessageById(client, userEmail, lastMessageId)
	if err != nil {
		fmt.Printf("Unable to get message id: %v", lastMessageId)
	} else {
		fmt.Println(lastMessage)
	}
	archivedMessage, err := gmail2.ArchiveMessage(client, userEmail, lastMessageId)
	if err != nil {
		fmt.Printf("Unable to archive message ID: %v", lastMessageId)
	} else {
		fmt.Println(archivedMessage)
	}



}
