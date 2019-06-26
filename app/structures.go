package main

import "time"

// Request response from vRO when you ask for a request
type Request struct {
	Type          string      `json:"@type"`
	ID            string      `json:"id"`
	IconID        string      `json:"iconId"`
	Version       int         `json:"version"`
	RequestNumber int         `json:"requestNumber"`
	State         string      `json:"state"`
	Description   interface{} `json:"description"`
	Reasons       interface{} `json:"reasons"`
	RequestedFor  string      `json:"requestedFor"`
	RequestedBy   string      `json:"requestedBy"`
	Organization  struct {
		TenantRef      string `json:"tenantRef"`
		TenantLabel    string `json:"tenantLabel"`
		SubtenantRef   string `json:"subtenantRef"`
		SubtenantLabel string `json:"subtenantLabel"`
	} `json:"organization"`
	RequestorEntitlementID string      `json:"requestorEntitlementId"`
	PreApprovalID          interface{} `json:"preApprovalId"`
	PostApprovalID         interface{} `json:"postApprovalId"`
	DateCreated            time.Time   `json:"dateCreated"`
	LastUpdated            time.Time   `json:"lastUpdated"`
	DateSubmitted          time.Time   `json:"dateSubmitted"`
	DateApproved           interface{} `json:"dateApproved"`
	DateCompleted          interface{} `json:"dateCompleted"`
	Quote                  struct {
		LeasePeriod    interface{} `json:"leasePeriod"`
		LeaseRate      interface{} `json:"leaseRate"`
		TotalLeaseCost interface{} `json:"totalLeaseCost"`
	} `json:"quote"`
	RequestCompletion struct {
		RequestCompletionState string      `json:"requestCompletionState"`
		CompletionDetails      string      `json:"completionDetails"`
		ResourceBindingIds     interface{} `json:"resourceBindingIds"`
	} `json:"requestCompletion"`
	RequestData struct {
		Entries []struct {
			Key   string      `json:"key"`
			Value interface{} `json:"value"`
		} `json:"entries"`
	} `json:"requestData"`
	RetriesRemaining         int         `json:"retriesRemaining"`
	RequestedItemName        string      `json:"requestedItemName"`
	RequestedItemDescription string      `json:"requestedItemDescription"`
	Components               interface{} `json:"components"`
	Successful               bool        `json:"successful"`
	Final                    bool        `json:"final"`
	StateName                string      `json:"stateName"`
	CatalogItemRef           struct {
		ID    string `json:"id"`
		Label string `json:"label"`
	} `json:"catalogItemRef"`
	CatalogItemProviderBinding struct {
		BindingID   string `json:"bindingId"`
		ProviderRef struct {
			ID    string `json:"id"`
			Label string `json:"label"`
		} `json:"providerRef"`
	} `json:"catalogItemProviderBinding"`
	ExecutionStatus string `json:"executionStatus"`
	WaitingStatus   string `json:"waitingStatus"`
	ApprovalStatus  string `json:"approvalStatus"`
	Phase           string `json:"phase"`
}

// CatalogItemInformations has information mendatory to request a catalog item on vRA
type CatalogItemInformations struct {
	CatalogItemID string `json:"catalogItemID"`
	SubtenantID   string `json:"subtenantID"`
	Requester     string `json:"requester"`
	Body          []byte `json:"body"`
}

// ResourceActionInformations has information mendatory to request a catalog item on vRA
type ResourceActionInformations struct {
	ActionID    string `json:"actionID"`
	SubtenantID string `json:"subtenantID"`
	Requester   string `json:"requester"`
	Body        []byte `json:"body"`
}

// FormDetail is vRA response when you ask for formDetail
type FormDetail struct {
	Layout struct {
		Pages []struct {
			ID    interface{} `json:"id"`
			Label interface{} `json:"label"`
			State struct {
				Dependencies []interface{} `json:"dependencies"`
				Facets       []interface{} `json:"facets"`
			} `json:"state"`
			Sections []struct {
				ID    interface{} `json:"id"`
				Label interface{} `json:"label"`
				State struct {
					Dependencies []interface{} `json:"dependencies"`
					Facets       []interface{} `json:"facets"`
				} `json:"state"`
				Rows []struct {
					Items []struct {
						Type             string      `json:"type"`
						Size             int         `json:"size"`
						ID               interface{} `json:"id"`
						ExtensionID      string      `json:"extensionId"`
						ExtensionPointID interface{} `json:"extensionPointId"`
						FieldPrefix      string      `json:"fieldPrefix"`
						State            struct {
							Dependencies []interface{} `json:"dependencies"`
							Facets       []interface{} `json:"facets"`
						} `json:"state"`
					} `json:"items"`
				} `json:"rows"`
			} `json:"sections"`
		} `json:"pages"`
	} `json:"layout"`
	Values        interface{} `json:"values"`
	FieldPrefixes interface{} `json:"fieldPrefixes"`
}

// VraEndpoint is the body structure used for token request
type VraEndpoint struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Tenant   string `json:"tenant"`
	Fqdn     string `json:"fqdn"`
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
