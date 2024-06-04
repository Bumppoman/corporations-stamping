package main

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"

	"github.com/koltyakov/gosip"
	strategy "github.com/koltyakov/gosip-sandbox/strategies/ondemand"
	"github.com/koltyakov/gosip/api"
)

// App struct
type StampService struct {}

type StampingItem struct {
	ID int `json:"Id"`
	CreationDate string `json:"CreationDate"`
	Selected bool `json:"Selected"`
	StagedforFiling bool `json:"StagedforFiling"`
	StampText string `json:"StampText"`
	SubmitterName string `json:"SubmitterName"`
}

// Download the unstamped attachment
func (srv *StampService) DownloadAttachment(id int) string {
	sp := getClient()

	// Load review item
	item := sp.Web().
		GetList("Lists/OathOfOfficeReviews1").
		Items().
		GetByID(id)

	// Load attachments
	attachments, _ := item.
		Attachments().
		Get()

	// Get the first attachment (unstamped item)
	pdfFilename := attachments.Data()[0].Data().FileName
	attachment, _ := item.Attachments().GetByName(pdfFilename).Download()

	// Return the Base64 encoded unstamped item
	return base64.StdEncoding.EncodeToString(attachment)
}

func (srv *StampService) LoadUnstamped() []StampingItem {
	sp := getClient()

	// Load unstamped review items
	listItems, _ := sp.Web().
		GetList("Lists/OathOfOfficeReviews1").
		Items().
		Select("Id,CreationDate,StagedforFiling,SubmitterName").
		Filter("StagedforFiling eq null and Filing/Determination eq 'Accepted'").
		Get()

	// Unmarshal the JSON into a Go struct
	items := []StampingItem{}
	json.Unmarshal(listItems.Normalized(), &items)

	// Return the list of unstamped review items
	return items
}

func (srv *StampService) SignIn() *api.UserInfo {
	// Get the SharePoint client and authentication configuration
	client, auth := getClientWithAuth()
	sp := api.NewSP(client)

	// Get the current user; if there is an error, clear the cookie cache and try again
	response, err := sp.Web().CurrentUser().Get()
	if err != nil {
		auth.CleanCookieCache()
		sp = api.NewSP(client)
		response, _ = sp.Web().CurrentUser().Get()
	}

	// Return the current user
	return response.Data()
}

// Upload the stamped attachment
func (srv *StampService) UploadStamped(id int, stamped string) error {
	// Decode the Base64 encoded stamped attachment
	pdfArray, _ := base64.StdEncoding.DecodeString(stamped)
	pdf := bytes.NewReader(pdfArray)

	// Get the review item
	sp := getClient()
	item := sp.Web().GetList("Lists/OathOfOfficeReviews1").Items().GetByID(id)

	// Remove unstamped attachment
	attachments, _ := item.Attachments().Get()
	attachment := attachments.Data()[0].Data().FileName
	err := item.Attachments().GetByName(attachment).Delete()
	if err != nil {
		err := item.Attachments().GetByName(attachment).Delete()
		if err != nil {
			return err
		}
	}

	// Add stamped attachment
	_, err = item.Attachments().Add("stamped.pdf", pdf)
	if err != nil {
		_, err = item.Attachments().Add("stamped.pdf", pdf)
		if err != nil {
			return err
		}
	}

	// Update `StagedforFiling` timestamp
	_, err = item.Update(
		[]byte(
			fmt.Sprintf(
				`{"StagedforFiling": "%s"}`,
				time.Now().Format(time.RFC3339),
			),
		),
	)

	return err
}

func getClient() *api.SP {
	client, _ := getClientWithAuth()
	return api.NewSP(client)
}

func getClientWithAuth() (*gosip.SPClient, *strategy.AuthCnfg) {
	auth := &strategy.AuthCnfg {
		SiteURL: "https://nysemail.sharepoint.com/sites/DOS/corp/Data",
	}

	client := &gosip.SPClient{AuthCnfg: auth}
	return client, auth
}
