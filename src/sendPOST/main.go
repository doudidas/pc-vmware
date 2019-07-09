package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

var url, token, body string
var reader io.Reader

func init() {
	if len(os.Args) < 2 {
		getUserInputs()
	} else {
		getUserArguments()
	}
}

func main() {

	// Create a new request using http
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(body)))

	if err != nil {
		KO("Failed to initiate POST Request: " + err.Error())
	}

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
		KO("Failed to send POST Request: " + err.Error())

	}

	body, err := ioutil.ReadAll(resp.Body)

	defer resp.Body.Close()

	if err != nil {
		KO(err.Error())
	}

	output := string(body)

	fmt.Println(resp.Header, output)
}

func getUserInputs() {
	fmt.Print("Please Provide the url (mendatory): ")
	_, err := fmt.Scanln(&url)
	if err != nil {
		KO(err.Error())
	}
	if url == "" {
		getUserInputs()
		return
	}
	fmt.Print("Please provide a bearer token if needed: ")
	_, err = fmt.Scanln(&token)
	if err != nil {
		KO(err.Error())
	}
	fmt.Print("Please provide a body if needed: ")
	_, err = fmt.Scanln(&body)
	if err != nil {
		KO(err.Error())
	}
	fmt.Println("url: ", url)
	fmt.Println("token: ", token)
	fmt.Println("body: ", body)
}

func getUserArguments() {
	url = os.Args[1]

	if len(os.Args) < 3 {
		token = ""
	} else {
		token = os.Args[2]
	}
	if len(os.Args) < 4 {
		body = ""
	} else {
		body = os.Args[3]
	}
}

// KO will generate a formatd error message
func KO(msg string) {
	tmp := ErrorWithInputs{
		Error: msg,
		URL:   url,
		Token: token,
		Body:  body,
	}
	output, _ := json.MarshalIndent(tmp, "", "   ")
	panic(string(output))
}

// ErrorWithInputs is the aggregation of error and inputs
type ErrorWithInputs struct {
	Error string `json:"error"`
	Token string `json:"token"`
	URL   string `json:"url"`
	Body  string `json:"body"`
}
