package troutdelete

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	troutclient "trout/client"
	troutstructs "trout/structs"
)

func DeleteUser(a, p, u, s string) {
	// retreive full user list
	users, err := getUsersList(a, p, s)
	if err != nil {
		fmt.Println("Error response:", err)
		return
	}

	// find the user ID for provided username
	var userID int
	userFound := false
	for _, user := range users {
		if user.Username == u {
			userID = user.ID
			userFound = true
			break
		}
	}

	if !userFound {
		fmt.Println("User not found:", u)
		return
	}

	deleteURL := "https://" + p + ":3333/api/users/" + strconv.Itoa(userID)

	client, err := troutclient.CreateHTTPClient(s)
	if err != nil {
		fmt.Println("Error creating HTTP client:", err)
	}

	req, err := http.NewRequest("DELETE", deleteURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", a)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var deleteResponse troutstructs.DeleteResponse

	if err := json.Unmarshal(body, &deleteResponse); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	if deleteResponse.Success {
		fmt.Printf("User '%s' (ID: %d) deleted successfully\n", u, userID)
	} else {
		fmt.Println("Error deleting user:", deleteResponse.Message)
	}
}

func getUsersList(a, p, s string) ([]troutstructs.User, error) {
	url := "https://" + p + ":3333/api/users/"

	client, err := troutclient.CreateHTTPClient(s)
	if err != nil {
		fmt.Println("Error creating HTTP client:", err)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", a)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var forbidden troutstructs.Forbidden
	if err := json.Unmarshal(body, &forbidden); err == nil && !forbidden.Success {
		return nil, fmt.Errorf("%s - must be 'admin' to delete users", forbidden.Message)
	}

	var users []troutstructs.User
	if err := json.Unmarshal(body, &users); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return users, nil
}

func DeleteCampaign(a, p, c, s string) {
	// retreive full campaign list
	campaigns, err := getCampaignList(a, p, s)
	if err != nil {
		fmt.Println("Error response:", err)
		return
	}

	// find the campaign ID for provided campaign name
	var campaignID int
	campaignFound := false
	for _, campaign := range campaigns {
		if campaign.Name == c {
			campaignID = campaign.ID
			campaignFound = true
			break
		}
	}

	if !campaignFound {
		fmt.Println("Campaign not found:", c)
		return
	}

	deleteURL := "https://" + p + ":3333/api/campaigns/" + strconv.Itoa(campaignID)

	client, err := troutclient.CreateHTTPClient(s)
	if err != nil {
		fmt.Println("Error creating HTTP client:", err)
	}

	req, err := http.NewRequest("DELETE", deleteURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", a)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var deleteResponse troutstructs.DeleteResponse

	if err := json.Unmarshal(body, &deleteResponse); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	if deleteResponse.Success {
		fmt.Printf("Campaign '%s' (ID: %d) deleted successfully\n", c, campaignID)
	} else {
		fmt.Println("Error deleting user:", deleteResponse.Message)
	}
}

func getCampaignList(a, p, s string) ([]troutstructs.Campaigns, error) {
	url := "https://" + p + ":3333/api/campaigns/"

	client, err := troutclient.CreateHTTPClient(s)
	if err != nil {
		fmt.Println("Error creating HTTP client:", err)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", a)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var forbidden troutstructs.Forbidden
	if err := json.Unmarshal(body, &forbidden); err == nil && !forbidden.Success {
		return nil, fmt.Errorf("%s", forbidden.Message)
	}

	var campaigns []troutstructs.Campaigns
	if err := json.Unmarshal(body, &campaigns); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return campaigns, nil
}

func DeleteSmtp(a, p, e, s string) {
	// retreive full sending profile list
	smtps, err := getSmtpList(a, p, s)
	if err != nil {
		fmt.Println("Error response:", err)
		return
	}

	// find the campaign ID for provided campaign name
	var smtpID int
	smtpFound := false
	for _, smtp := range smtps {
		if smtp.Name == e {
			smtpID = smtp.ID
			smtpFound = true
			break
		}
	}

	if !smtpFound {
		fmt.Println("Campaign not found:", e)
		return
	}

	deleteURL := "https://" + p + ":3333/api/smtp/" + strconv.Itoa(smtpID)

	client, err := troutclient.CreateHTTPClient(s)
	if err != nil {
		fmt.Println("Error creating HTTP client:", err)
	}

	req, err := http.NewRequest("DELETE", deleteURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", a)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var deleteResponse troutstructs.DeleteResponse

	if err := json.Unmarshal(body, &deleteResponse); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	if deleteResponse.Success {
		fmt.Printf("Sending profile '%s' (ID: %d) deleted successfully\n", e, smtpID)
	} else {
		fmt.Println("Error deleting user:", deleteResponse.Message)
	}
}

func getSmtpList(a, p, s string) ([]troutstructs.SMTP, error) {
	url := "https://" + p + ":3333/api/smtp/"

	client, err := troutclient.CreateHTTPClient(s)
	if err != nil {
		fmt.Println("Error creating HTTP client:", err)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Authorization", a)

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	var forbidden troutstructs.Forbidden
	if err := json.Unmarshal(body, &forbidden); err == nil && !forbidden.Success {
		return nil, fmt.Errorf("%s", forbidden.Message)
	}

	var smtps []troutstructs.SMTP
	if err := json.Unmarshal(body, &smtps); err != nil {
		return nil, fmt.Errorf("error parsing JSON: %v", err)
	}

	return smtps, nil
}
