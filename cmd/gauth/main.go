package main

import (
	"encoding/base64"
	"fmt"
	"github.com/negronjl/libGoogle/pkg/gauth"
	"log"
	"os"
	"flag"
)

func main() {
	cjson := flag.String("cjson", "", " Location of credentials.json file")
	scopes := flag.String("scopes", "", " Locations of the scopes.json file")
	flag.Parse()

	if *cjson == "" || *scopes == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	token, err := gauth.GetToken(*cjson, *scopes)
	if err != nil {
		log.Fatalf("Unable to get token: %v", err)
	}

	fmt.Printf("Token: [%v]\n", base64.StdEncoding.EncodeToString(token))
}
