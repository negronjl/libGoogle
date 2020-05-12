package gauth

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/negronjl/libGoogle/internal/common"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io/ioutil"
)

func LoadCredentialsJSON(fileName string) (common.GoogleAppCredentials, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil{
		return common.GoogleAppCredentials{}, err
	}
	var returnValue common.GoogleAppCredentials
	err = json.Unmarshal(b, &returnValue)
	if err != nil {
		return common.GoogleAppCredentials{}, err
	}
	return returnValue, nil
}

func LoadScopes(fileName string) (common.GoogleScopes, error) {
	b, err := ioutil.ReadFile(fileName)
	if err != nil {
		return common.GoogleScopes{}, err
	}
	var returnValue common.GoogleScopes
	err = json.Unmarshal(b, &returnValue)
	if err != nil {
		return common.GoogleScopes{}, err
	}
	return returnValue, nil
}

func GetClientConfig(credentials common.GoogleAppCredentials, scopes common.GoogleScopes) (*oauth2.Config, error) {
	c, err := json.Marshal(credentials)
	if err != nil {
		return nil, err
	}
	config, err := google.ConfigFromJSON(c, scopes.Scopes...)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func GetTokenFromWeb(config *oauth2.Config) (*oauth2.Token, error) {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		return nil, err
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		return nil, err
	}
	return tok, nil
}

func GetTokenFromB64String(b64Token string) (*oauth2.Token, error) {
	if b64Token == "" {
		return nil, errors.New("Empty b64Token ")
	}
	clearToken, err := base64.StdEncoding.DecodeString(b64Token)
	if err != nil {
		return nil, err
	}
	var token oauth2.Token
	err = json.NewDecoder(bytes.NewReader(clearToken)).Decode(&token)
	if err != nil {
		return nil, err
	}
	return &token, nil
}

