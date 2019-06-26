package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/fatih/color"
)

var token string
var endpoint VraEndpoint
var debug = true

func main() {

	// var r Request
	color.Green("START - Demo vRO get action output")
	initConf()
	_ = requestCatalogItem("provisionVM")

	// r = requestAction("destroyVM", r.CatalogItemRef.ID)
}

func requestCatalogItem(configFileName string) Request {
	var output []byte
	var r Request
	var executionStatus string
	info := getcatalogInformationsFromFile("../resources/" + configFileName + ".json")
	if info.Body == nil {
		info.Body = getCatalogTemplate(info.CatalogItemID)
		color.Green("catalogItem form done !")
	}
	color.Cyan("POST Catalog Item Request")

	url := "https://" + endpoint.Fqdn + "/catalog-service/api/consumer/entitledCatalogItems/" + info.CatalogItemID + "/requests?businessGroupId=" + info.SubtenantID + "&requestedFor=" + info.Requester
	_, header := POST(url, token, info.Body)
	color.Blue("location:" + header.Get("Location"))

	output = GET(header.Get("Location"), token)
	json.Unmarshal(output, &r)
	executionStatus = r.ExecutionStatus
	color.Cyan("***** WAIT FOR Catalog Item COMPLETION ******")
	for executionStatus != "STOPPED" {
		fmt.Println(r.Phase)
		time.Sleep(5 * time.Second)
		output = GET(header.Get("Location"), token)
		json.Unmarshal(output, &r)
		executionStatus = r.ExecutionStatus
	}
	if r.Phase == "FAILED" {
		color.Red("FAILED !")
		color.Red(r.RequestCompletion.CompletionDetails)
		panic("Failed to request CatalogItem !" + r.RequestCompletion.CompletionDetails)
	}
	color.Green("DONE !")
	color.Cyan("GET FORM OUTPUT")
	result := GET(header.Get("Location")+"/forms/details", token)
	color.Green("DONE !")
	color.Cyan("SHOW RESULT")
	fmt.Println(string(result))
	color.Green("DONE")
	return r
}

func requestAction(configFileName, resourceID string) Request {
	var output []byte
	var r Request
	var executionStatus string
	var formDetail FormDetail
	info := getResourceActionInformationsFromFile("../resources/" + configFileName + ".json")
	if info.Body == nil {
		info.Body = getActionTemplate(info, resourceID)
		color.Green("action form done !")
	}

	color.Cyan("POST Request Action")
	url := "https://" + endpoint.Fqdn + "/catalog-service/api/consumer/resources/" + resourceID + "/actions/" + info.ActionID + "/requests"
	_, header := POST(url, token, info.Body)
	output = GET(header.Get("Location"), token)
	json.Unmarshal(output, &r)
	executionStatus = r.ExecutionStatus
	color.Cyan("***** WAIT FOR ACTION COMPLETION ******")
	for executionStatus != "STOPPED" {
		fmt.Println(executionStatus)
		time.Sleep(5 * time.Second)
		output = GET(header.Get("Location"), token)
		json.Unmarshal(output, &r)
		executionStatus = r.ExecutionStatus
	}
	if r.Phase == "FAILED" {
		color.Red("FAILED !")
		color.Red(r.RequestCompletion.CompletionDetails)
		panic("Failed to request CatalogItem !" + r.RequestCompletion.CompletionDetails)
	}
	color.Green("DONE !")
	color.Cyan("GET FORM OUTPUT")
	result := GET(header.Get("Location")+"/forms/details", token)
	json.Unmarshal(result, &formDetail)

	color.Green("DONE !")
	color.Cyan("SHOW RESULT")
	fmt.Println(formDetail.Values)
	color.Green("DONE")
	return r
}

func getCatalogTemplate(catalogID string) []byte {
	fmt.Println("***** Looking for catalog template ******")
	url := "https://" + endpoint.Fqdn + "/catalog-service/api/consumer/entitledCatalogItems/" + catalogID + "/requests/template"
	output := GET(url, token)

	return output
}

func getActionTemplate(info ResourceActionInformations, resourceID string) []byte {
	fmt.Println("***** Looking for action template ******")
	url := "https://" + endpoint.Fqdn + "/catalog-service/api/consumer/resources/" + resourceID + "/actions/" + info.ActionID + "/requests/template"
	output := GET(url, token)
	return output
}

func initConf() {
	color.Green("START INT")
	endpoint = getVRAEndpoint()
	token = getToken(endpoint)
	color.Green("INIT DONE !")
}
