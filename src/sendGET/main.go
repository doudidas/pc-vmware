package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var url, token string
var outputError OutputError

func init() {

	if len(os.Args) < 2 {
		getUserInputs()
	} else {
		getUserArguments()
	}

}

// GET allows to do get call to vRA
func main() {

	// Create a new request using http
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		KO("Failed to initiate GET Request: " + err.Error())
	}
	if token != "" {
		// Create a Bearer string by appending string access token
		var bearer = "Bearer " + token
		// Add authorization header to the req
		req.Header.Add("Authorization", bearer)
	}

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	// Send req using http Client
	client := &http.Client{Transport: tr}
	resp, err := client.Do(req)

	if err != nil {
		KO("Failed to execute GET Request: " + err.Error())
	}

	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	var bodyAsError APIError
	json.Unmarshal(body, &bodyAsError)
	if err != nil {
		KO(string(body))
	}

	fmt.Println(string(body))
}

// KO will generate a formatd error message
func KO(msg string) {
	tmp := ErrorWithInputs{
		Error: msg,
		URL:   url,
		Token: token,
	}
	output, _ := json.MarshalIndent(tmp, "", "   ")
	panic(string(output))
}

func getUserInputs() {
	for url == "" {
		fmt.Println("Please Provide the API url (mendatory): ")
		nb, err := fmt.Scanln(&url)
		if err != nil && nb != 0 {
			KO(err.Error())
		}
	}

	fmt.Print("Please provide a bearer token if needed: ")
	nb, err := fmt.Scanln(&token)
	if nb == 0 {
		token = ""
	}
	if err != nil {
		defer KO("Failed to get user inputs " + err.Error())
	}
}

func getUserArguments() {
	url = os.Args[1]

	if len(os.Args) < 3 {
		token = ""
	} else {
		token = os.Args[2]
	}
}

// ErrorWithInputs is the aggregation of error and inputs
type ErrorWithInputs struct {
	Error string `json:"error"`
	Token string `json:"token"`
	URL   string `json:"url"`
}

// APIError is the format of error response from vRA
type APIError struct {
	Errors []struct {
		Code          int         `json:"code"`
		Source        interface{} `json:"source"`
		Message       string      `json:"message"`
		SystemMessage string      `json:"systemMessage"`
		MoreInfoURL   interface{} `json:"moreInfoUrl"`
	} `json:"errors"`
}
