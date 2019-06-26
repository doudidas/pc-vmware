package main

import (
	"encoding/json"
	"log"
	"os"
)

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
func getResourceActionInformationsFromFile(fileName string) ResourceActionInformations {
	var output ResourceActionInformations
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&output)
	return output
}
func getVRAEndpoint() VraEndpoint {
	var v VraEndpoint
	file, err := os.Open("../resources/vra-endpoint.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	json.NewDecoder(file).Decode(&v)
	return v
}
