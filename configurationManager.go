package main

import (
	"encoding/json"
	"log"
	"os"
)

func getTokenCredantialFromSecretFile() TokenCredantial {
	var output TokenCredantial
	file, err := os.Open("secret.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&output)
	return output
}

func getcatalogInformationsFromFile(fileName string) CatalogItemInformations {
	var output CatalogItemInformations
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&output)
	return output
}
