package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/fatih/color"
)

var token string

var debug = false

func main() {
	color.Green("START - Demo vRO get action output")
	token := getToken()
	info := getcatalogInformationsFromFile("provisionVM.json")
	body := getCatalogTemplate(token, info.CatalogItemID)
	color.Green("catalogItem form done !")
	requestCatalogItem(token, info, body)
}

func requestCatalogItem(token string, info CatalogItemInformations, body []byte) {
	var output []byte
	var i Request
	var executionStatus string
	color.Cyan("POST Catalog Item Request")

	url := "https://cava-n-80-154.eng.vmware.com/catalog-service/api/consumer/entitledCatalogItems/" + info.CatalogItemID + "/requests?businessGroupId=" + info.SubtenantID + "&requestedFor=" + info.Requester
	_, header := POST(url, token, body)
	color.Blue("location:" + header.Get("Location"))

	output = GET(header.Get("Location"), token)
	json.Unmarshal(output, &i)
	executionStatus = i.ExecutionStatus
	color.Cyan("***** WAIT FOR Catalog Item COMPLETION ******")
	for executionStatus != "STOPPED" {
		fmt.Println(i.Phase)
		time.Sleep(5 * time.Second)
		output = GET(header.Get("Location"), token)
		json.Unmarshal(output, &i)
		executionStatus = i.ExecutionStatus
	}
	if i.Phase == "FAILED" {
		color.Red("FAILED !")
		color.Red(i.RequestCompletion.CompletionDetails)
		panic("Failed to request CatalogItem !" + i.RequestCompletion.CompletionDetails)
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
	var formDetail FormDetail
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
	json.Unmarshal(result, &formDetail)

	color.Green("DONE !")
	color.Cyan("SHOW RESULT")
	fmt.Println(formDetail.Values)
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
