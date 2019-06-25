package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/fatih/color"
)

// GET allows to do get call to vRA
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
	if debug {
		fmt.Println(string(output))
	}
	return output
}

// POST allows to do post call to vRA
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
	if debug {
		fmt.Println(string(output))
	}
	return output, resp.Header
}

// getToken will read the token file and will use it to do any API call
func getToken() string {
	url := "https://cava-n-80-154.eng.vmware.com/identity/api/tokens"
	color.Cyan("GET TOKEN")
	color.Magenta("url: %s", url)
	message := getTokenCredantialFromSecretFile()

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
