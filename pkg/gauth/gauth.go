package gauth

import (
	"encoding/json"
	"github.com/negronjl/libGoogle/internal/gauth"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

func GetToken(credentialsJsonFileName string, scopesFileName string) ([]byte, error) {
	cjson, err := gauth.LoadCredentialsJSON(credentialsJsonFileName)
	if err != nil {
		return nil, err
	}
	scopes, err := gauth.LoadScopes(scopesFileName)
	if err != nil {
		return nil, err
	}
	clientConfig, err := gauth.GetClientConfig(cjson, scopes)
	if err != nil {
		return nil, err
	}
	token, err := gauth.GetTokenFromWeb(clientConfig)
	if err != nil {
		return nil, err
	}
	var returnValue []byte
	returnValue, err = json.Marshal(token)
	if err != nil {
		return nil, err
	}
	return returnValue, nil
}

func GetClient(credentialsJsonFileName string, scopesFileName string, b64Token string) *http.Client {
	if credentialsJsonFileName == "" || scopesFileName == "" || b64Token == "" {
		return nil
	}
	cjson, err := gauth.LoadCredentialsJSON(credentialsJsonFileName)
	if err != nil {
		log.Printf("Error Loading Credentials JSON: %v", err)
		return nil
	}
	scopes, err := gauth.LoadScopes(scopesFileName)
	if err != nil {
		log.Printf("Error Loading Scopes JSON file: %v", err)
		return nil
	}
	clientConfig, err := gauth.GetClientConfig(cjson, scopes)
	if err != nil {
		log.Printf("Error getting client configuration: %v", err)
		return nil
	}
	token, err := gauth.GetTokenFromB64String(b64Token)
	if err != nil {
		log.Printf("Error Getting token from B64 string: %v", err)
		return nil
	}
	return clientConfig.Client(context.Background(), token)
}
