package troutlist

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
	troutclient "trout/client"
	troutstructs "trout/structs"

	"github.com/fatih/color"
)

func GetTemplates(directory string, typeName string, fileExtension string) {
	green := color.New(color.FgGreen).PrintfFunc()
	green("\n%s templates available:\n", typeName)

	files, err := os.ReadDir(directory)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == fileExtension {
			filePath := filepath.Join(directory, file.Name())
			fileData, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}

			if typeName == "Email" {
				var template troutstructs.EmailTemplate
				if err := json.Unmarshal(fileData, &template); err != nil {
					fmt.Println("Error unmarshalling JSON:", err)
					continue
				}

				fmt.Println(template.Name)
			} else if typeName == "Landing page" {
				var template troutstructs.PageTemplate
				if err := json.Unmarshal(fileData, &template); err != nil {
					fmt.Println("Error unmarshalling JSON:", err)
					continue
				}

				fmt.Println(template.Name)
			}

		}
	}
}

func GetEmails(a, p, s string) {
	url := "https://" + p + ":3333/api/templates/"

	client, err := troutclient.CreateHTTPClient(s)
	if err != nil {
		fmt.Println("Error creating HTTP client:", err)
	}

	req, err := http.NewRequest("GET", url, nil)
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

	var templates []troutstructs.EmailTemplate
	if err := json.Unmarshal(body, &templates); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("Email Templates:")
	for _, template := range templates {
		fmt.Printf("ID: %d\nName: %s\n----------\n", template.ID, template.Name)
	}
}

func GetPages(a, p, s string) {
	url := "https://" + p + ":3333/api/pages/"

	client, err := troutclient.CreateHTTPClient(s)
	if err != nil {
		fmt.Println("Error creating HTTP client:", err)
	}

	req, err := http.NewRequest("GET", url, nil)
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

	var templates []troutstructs.PageTemplate
	if err := json.Unmarshal(body, &templates); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("Landing Page Templates:")
	for _, template := range templates {
		fmt.Printf("ID: %d\nName: %s\n----------\n", template.ID, template.Name)
	}
}

func GetSmtp(a, p, s string) {
	url := "https://" + p + ":3333/api/smtp/"

	client, err := troutclient.CreateHTTPClient(s)
	if err != nil {
		fmt.Println("Error creating HTTP client:", err)
	}

	req, err := http.NewRequest("GET", url, nil)
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

	var smtps []troutstructs.SMTP
	if err := json.Unmarshal(body, &smtps); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("Sending profiles:")
	for _, smtp := range smtps {
		fmt.Printf("ID: %d\nName: %s\n----------\n", smtp.ID, smtp.Name)
	}
}

func GetCampaigns(a, p, s string) {
	url := "https://" + p + ":3333/api/campaigns/"

	client, err := troutclient.CreateHTTPClient(s)
	if err != nil {
		fmt.Println("Error creating HTTP client:", err)
	}

	req, err := http.NewRequest("GET", url, nil)
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

	var campaigns []troutstructs.Campaigns
	if err := json.Unmarshal(body, &campaigns); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("Campaigns:")
	for _, campaign := range campaigns {
		launchDate, err := time.Parse(time.RFC3339, campaign.LaunchDate)
		if err != nil {
			fmt.Println("Error parsing launch date:", err)
			return
		}

		formattedLaunchDate := launchDate.Format("2006-01-02")

		fmt.Printf("ID: %d\nName: %s\nLaunch Date: %s\n----------\n", campaign.ID, campaign.Name, formattedLaunchDate)
	}
}

func GetUsers(a, p, s string) {
	url := "https://" + p + ":3333/api/users/"

	client, err := troutclient.CreateHTTPClient(s)
	if err != nil {
		fmt.Println("Error creating HTTP client:", err)
	}

	req, err := http.NewRequest("GET", url, nil)
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

	var forbidden troutstructs.Forbidden
	if err := json.Unmarshal(body, &forbidden); err == nil && !forbidden.Success {
		fmt.Println("User must be 'admin' to list users")
		return
	}

	var users []troutstructs.User
	if err := json.Unmarshal(body, &users); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fmt.Println("Users:")
	for _, user := range users {
		role := user.Role
		fmt.Printf("ID: %d, Role: %s, API Key: %s, Username: %s \n", user.ID, role.Name, user.ApiKey, user.Username)
	}
}
