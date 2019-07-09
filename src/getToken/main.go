package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

var fqdn, username, password, tenant, url string

func init() {
	if len(os.Args) < 2 {
		getUserInputs()
	} else {
		getUserArguments()
	}
}

// getToken will read the token file and will use it to do any API call
func main() {

	url := "https://" + fqdn + "/identity/api/tokens"

	// Prepare the body to send
	body := map[string]string{
		"username": username,
		"password": password,
		"tenant":   tenant,
	}
	// Convert body as bytes
	bodyAsBytes, err := json.Marshal(body)
	if err != nil {
		log.Fatalln(err)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	http.DefaultClient.Transport = tr
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(bodyAsBytes))
	if err != nil {
		log.Fatalln(err)
	}
	var result map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result["id"].(string))
}

func getUserInputs() {
	for fqdn == "" {
		fmt.Println("Please Provide the vra fqdn (mendatory): ")
		nb, err := fmt.Scanln(&fqdn)
		if err != nil && nb != 0 {
			KO(err.Error())
		}
	}
	for tenant == "" {
		fmt.Println("Please Provide the tenant (mendatory): ")
		nb, err := fmt.Scanln(&tenant)
		if err != nil && nb != 0 {
			KO(err.Error())
		}
	}
	for username == "" {
		fmt.Println("Please Provide the full username (Ex: api-user@vsphere.local) (mendatory): ")
		nb, err := fmt.Scanln(&username)
		if err != nil && nb != 0 {
			KO(err.Error())
		}
	}
	for password == "" {
		fmt.Println("Please Provide the pasword (mendatory): ")
		nb, err := fmt.Scanln(&password)
		if err != nil && nb != 0 {
			KO(err.Error())
		}
	}
}

func getUserArguments() {
	fqdn = os.Args[1]
	tenant = os.Args[2]
	username = os.Args[2]
	password = os.Args[4]
}

// ErrorWithInputs is the aggregation of error and inputs
type ErrorWithInputs struct {
	Error    string `json:"error"`
	Token    string `json:"token"`
	Fqdn     string `json:"fqdn"`
	Username string `json:"username"`
	Password string `json:"password"`
	Tenant   string `json:"tenant"`
}

// KO will generate a formatd error message
func KO(msg string) {
	tmp := ErrorWithInputs{
		Error:    msg,
		Fqdn:     fqdn,
		Username: username,
		Password: password,
		Tenant:   tenant,
	}
	output, _ := json.MarshalIndent(tmp, "", "   ")
	panic(string(output))
}
