package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/fatih/color"
)

var token string

func main() {
	color.Green("START - Demo vRO get action output")
	token = getToken()
	body := getCatalogTemplate(token, "8b0c1cca-2678-470d-ab27-05f7c8a9fc21")
	requestCatalogItem(token, "8b0c1cca-2678-470d-ab27-05f7c8a9fc21", body)
	// requestAction(token, "9c32adef-f768-4faa-a934-a2356d0fe3fd", "1fc45d5d-17d6-4a71-8286-acf6e19a8228", body)
}

func requestCatalogItem(token, catalogID string, body []byte) {
	var output []byte
	var i map[string]string
	var executionStatus string
	color.Cyan("POST Request Action")
	var subtenantID = "b3a2dc51-5267-410c-87f3-ddfefc1645a7"
	var username = "etopin@vsphere.local"
	url := "https://cava-n-80-154.eng.vmware.com/catalog-service/api/consumer/entitledCatalogItems/" + catalogID + "/requests?businessGroupId=" + subtenantID + "&requestedFor=" + username
	_, header := POST(url, token, body)
	output = GET(header.Get("Location"), token)
	json.Unmarshal(output, &i)
	executionStatus = i["executionStatus"]
	color.Cyan("***** WAIT FOR ACTION COMPLETION ******")
	for executionStatus != "STOPPED" {
		fmt.Println(executionStatus)
		time.Sleep(5 * time.Second)
		output = GET(header.Get("Location"), token)
		json.Unmarshal(output, &i)
		executionStatus = i["executionStatus"]
	}
	color.Green("DONE !")
	color.Cyan("GET FORM OUTPUT")
	result := GET(header.Get("Location")+"/forms/details", token)
	color.Green("DONE !")
	color.Cyan("SHOW RESULT")
	fmt.Println(string(result))
	color.Green("DONE")
}

func requestAction(token, resourceID, actionID string, body []byte) {
	var output []byte
	var i map[string]string
	var executionStatus string
	color.Cyan("POST Request Action")
	url := "https://cava-n-80-154.eng.vmware.com/catalog-service/api/consumer/resources/" + resourceID + "/actions/" + actionID + "/requests"
	_, header := POST(url, token, body)
	output = GET(header.Get("Location"), token)
	json.Unmarshal(output, &i)
	executionStatus = i["executionStatus"]
	color.Cyan("***** WAIT FOR ACTION COMPLETION ******")
	for executionStatus != "STOPPED" {
		fmt.Println(executionStatus)
		time.Sleep(5 * time.Second)
		output = GET(header.Get("Location"), token)
		json.Unmarshal(output, &i)
		executionStatus = i["executionStatus"]
	}
	color.Green("DONE !")
	color.Cyan("GET FORM OUTPUT")
	result := GET(header.Get("Location")+"/forms/details", token)
	color.Green("DONE !")
	color.Cyan("SHOW RESULT")
	fmt.Println(string(result))
	color.Green("DONE")
}

func getCatalogTemplate(token, catalogID string) []byte {
	fmt.Println("***** Looking for catalog template ******")
	url := "https://cava-n-80-154.eng.vmware.com/catalog-service/api/consumer/entitledCatalogItems/" + catalogID + "/requests/template"
	output := GET(url, token)
	return output
}

func getActionTemplate(token, ressourceID, actionID string) []byte {
	fmt.Println("***** Looking for action template ******")
	url := "https://cava-n-80-154.eng.vmware.com/catalog-service/api/consumer/resources/" + ressourceID + "/actions/" + actionID + "/requests/template"
	output := GET(url, token)
	return output
}

func getToken() string {
	url := "https://cava-n-80-154.eng.vmware.com/identity/api/tokens"
	color.Cyan("GET TOKEN")
	color.Magenta("url: %s", url)
	message := map[string]interface{}{
		"username": "etopin@vsphere.local",
		"password": "VMware1!",
		"tenant":   "vsphere.local",
	}
	bytesRepresentation, err := json.Marshal(message)
	if err != nil {
		log.Fatalln(err)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	http.DefaultClient.Transport = tr
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bytesRepresentation))
	if err != nil {
		log.Fatalln(err)
	}
	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println("Token: ", result["id"].(string))
	return result["id"].(string)
}

func GET(url, token string) []byte {
	color.Yellow("GET")
	color.Magenta("url: %s", url)
	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	if token != "" {
		// Create a Bearer string by appending string access token
		var bearer = "Bearer " + token
		// add authorization header to the req
		req.Header.Add("Authorization", bearer)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Send req using http Client
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	output := []byte(body)
	fmt.Println(string(output))
	return output
}

func POST(url, token string, bytesRepresentation []byte) ([]byte, http.Header) {
	color.Yellow("POST")
	color.Magenta("url: %s", url)
	// Create a new request using http
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(bytesRepresentation))

	if token != "" {
		// Create a Bearer string by appending string access token
		var bearer = "Bearer " + token
		// add authorization header to the req
		req.Header.Add("Authorization", bearer)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Send req using http Client
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERRO] -", err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()
	output := []byte(body)
	fmt.Println(string(output))
	return output, resp.Header
}
