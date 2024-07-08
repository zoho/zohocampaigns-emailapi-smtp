package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Recipient struct {
	Address        string            `json:"address"`
	Name           string            `json:"name"`
	AdditionalData map[string]string `json:"additional_data"`
	MergeData      map[string]string `json:"merge_data"`
}

type Content struct {
	Subject string `json:"subject"`
	HTML    string `json:"html"`
	Text    string `json:"text"`
	From    struct {
		Address string `json:"address"`
		Name    string `json:"name"`
	} `json:"from"`
}

type Payload struct {
	CampaignName string      `json:"campaign_name"`
	Recipients   []Recipient `json:"recipients"`
	Content      Content     `json:"content"`
}

func main() {
	accessToken := "1000.****************************"

	// Recipients
	recipients := []Recipient{
		{
			Address: "lucy@example.campaigns.zoho.com",
			Name:    "Aaron Fletcher",
			AdditionalData: map[string]string{
				"phone":   "+919876543210",
				"country": "IN",
			},
			MergeData: map[string]string{
				"first_name": "Aaron",
			},
		},
	}

	// Content
	content := Content{
		Subject: "My first mail using Zoho Campaigns Email API HTTP",
		HTML:    "<html><body>Welcome $[first_name|Customer]$!<br>Summer Hot Savings, You Don't Want to Miss</body></html>",
		Text:    "Welcome $[first_name|Customer]$! Summer Hot Savings, You Don’t Want to Miss",
	}
	content.From.Address = "aron@marketing.campaigns.zoho.com"
	content.From.Name = "Aron Fletcher"

	// Payload
	payload := Payload{
		CampaignName: "hello_customer",
		Recipients:   recipients,
		Content:      content,
	}

	// Encode payload to JSON
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error marshalling payload: %v\n", err)
		return
	}

	// Set up HTTP request
	url := "https://campaigns.zoho.com/emailapi/v1/transmission"
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Authorization", "Zoho-oauthtoken "+accessToken)

	client := &http.Client{
		Timeout: 60 * time.Second,
	}

	// Send HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error sending request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	if resp.StatusCode >= 200 && resp.StatusCode <= 299 {
		fmt.Printf("Response: %s\n", body)
	} else {
		fmt.Printf("Error: %s\n", body)
	}
}
