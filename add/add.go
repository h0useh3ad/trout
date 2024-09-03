package troutadd

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
	troutclient "trout/client"
)

func AddEmails(a, p, s string) {
	url := "https://" + p + ":3333/api/templates/"

	dirPath := "add/templates/emails"
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	allSuccess := true

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".emailtemplate" {
			filePath := filepath.Join(dirPath, file.Name())
			fileData, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}

			var data map[string]interface{}
			if err := json.Unmarshal(fileData, &data); err != nil {
				fmt.Println("Error unmarshalling JSON:", err)
				fmt.Println("Failed JSON data:", string(fileData))
				continue
			}

			j, err := json.Marshal(data)
			if err != nil {
				fmt.Println("Error marshalling JSON:", err)
				continue
			}

			client, err := troutclient.CreateHTTPClient(s)
			if err != nil {
				fmt.Println("Error creating HTTP client:", err)
				continue
			}

			req, err := http.NewRequest("POST", url, strings.NewReader(string(j)))
			if err != nil {
				fmt.Println("Error creating request:", err)
				continue
			}

			req.Header.Set("Authorization", a)
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Error sending request:", err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusConflict {
				fmt.Printf("Template '%s' already present.\n", data["name"])
				allSuccess = false
				continue
			} else if resp.StatusCode != http.StatusCreated {
				fmt.Println("Error: Email template was not added successfully. Status code:", resp.StatusCode)
				allSuccess = false
			}
		}
	}

	if allSuccess {
		fmt.Println("All email templates added successfully.")
	}
}

func AddEmailContains(a, p, c, s string) {
	url := "https://" + p + ":3333/api/templates/"

	dirPath := "add/templates/emails"
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	found := false

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".emailtemplate" {
			filePath := filepath.Join(dirPath, file.Name())
			fileData, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}

			var data map[string]interface{}
			if err := json.Unmarshal(fileData, &data); err != nil {
				fmt.Println("Error unmarshalling JSON:", err)
				fmt.Println("Failed JSON data:", string(fileData))
				return
			}

			j, err := json.Marshal(data)
			if err != nil {
				fmt.Println("Error marshalling JSON:", err)
				return
			}

			if strings.Contains(strings.ToLower(string(j)), strings.ToLower(c)) {
				found = true
				client, err := troutclient.CreateHTTPClient(s)
				if err != nil {
					fmt.Println("Error creating HTTP client:", err)
					continue
				}

				req, err := http.NewRequest("POST", url, strings.NewReader(string(j)))
				if err != nil {
					fmt.Println("Error creating request:", err)
					continue
				}

				req.Header.Set("Authorization", a)
				req.Header.Set("Content-Type", "application/json")

				resp, err := client.Do(req)
				if err != nil {
					fmt.Println("Error sending request:", err)
					continue
				}
				defer resp.Body.Close()

				if resp.StatusCode == http.StatusConflict {
					fmt.Printf("Template '%s' already present.\n", data["name"])
					continue
				} else if resp.StatusCode != http.StatusCreated {
					fmt.Println("Error: Email template was not added successfully. Status code:", resp.StatusCode)
					continue
				}

				fmt.Printf("Template '%s' added successfully.\n", data["name"])
			}

		}
	}
	if !found {
		fmt.Printf("No template names contain '%s'\n", s)
	}
}

func AddPages(a, p, s string) {
	url := "https://" + p + ":3333/api/pages/"

	dirPath := "add/templates/pages"
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	allSuccess := true

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".pagetemplate" {
			filePath := filepath.Join(dirPath, file.Name())
			fileData, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}

			var data map[string]interface{}
			if err := json.Unmarshal(fileData, &data); err != nil {
				fmt.Println("Error unmarshalling JSON:", err)
				fmt.Println("Failed JSON data:", string(fileData))
				continue
			}

			j, err := json.Marshal(data)
			if err != nil {
				fmt.Println("Error marshalling JSON:", err)
				continue
			}

			client, err := troutclient.CreateHTTPClient(s)
			if err != nil {
				fmt.Println("Error creating HTTP client:", err)

			}

			req, err := http.NewRequest("POST", url, strings.NewReader(string(j)))
			if err != nil {
				fmt.Println("Error creating request:", err)
				continue
			}

			req.Header.Set("Authorization", a)
			req.Header.Set("Content-Type", "application/json")

			resp, err := client.Do(req)
			if err != nil {
				fmt.Println("Error sending request:", err)
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusConflict {
				fmt.Printf("Template '%s' already present.\n", data["name"])
				allSuccess = false
				continue
			} else if resp.StatusCode != http.StatusCreated {
				fmt.Println("Error: Email template was not added successfully. Status code:", resp.StatusCode)
				allSuccess = false
			}
		}
	}

	if allSuccess {
		fmt.Println("All landing pages added successfully.")
	}
}

func AddPageContains(a, p, c, s string) {
	url := "https://" + p + ":3333/api/pages/"

	dirPath := "add/templates/pages"
	files, err := os.ReadDir(dirPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	found := false

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".pagetemplate" {
			filePath := filepath.Join(dirPath, file.Name())
			fileData, err := os.ReadFile(filePath)
			if err != nil {
				fmt.Println("Error reading file:", err)
				continue
			}

			var data map[string]interface{}
			if err := json.Unmarshal(fileData, &data); err != nil {
				fmt.Println("Error unmarshalling JSON:", err)
				fmt.Println("Failed JSON data:", string(fileData))
				return
			}

			j, err := json.Marshal(data)
			if err != nil {
				fmt.Println("Error marshalling JSON:", err)
				return
			}

			if strings.Contains(strings.ToLower(string(j)), strings.ToLower(c)) {
				found = true
				client, err := troutclient.CreateHTTPClient(s)
				if err != nil {
					fmt.Println("Error creating HTTP client:", err)
					continue
				}

				req, err := http.NewRequest("POST", url, strings.NewReader(string(j)))
				if err != nil {
					fmt.Println("Error creating request:", err)
					continue
				}

				req.Header.Set("Authorization", a)
				req.Header.Set("Content-Type", "application/json")

				resp, err := client.Do(req)
				if err != nil {
					fmt.Println("Error sending request:", err)
					continue
				}
				defer resp.Body.Close()

				if resp.StatusCode == http.StatusConflict {
					fmt.Printf("Template '%s' already present.\n", data["name"])
					continue
				} else if resp.StatusCode != http.StatusCreated {
					fmt.Println("Error: Landing page template was not added successfully. Status code:", resp.StatusCode)
					continue
				}

				fmt.Printf("Template '%s' added successfully.\n", data["name"])
			}

		}
	}
	if !found {
		fmt.Printf("No template names contain '%s'\n", c)
	}
}

func AddUser(a, p, u, role, s string) {
	password := generatePassword()
	url := "https://" + p + ":3333/api/users/"
	data := map[string]string{
		"username": u,
		"password": password,
		"role":     role,
	}

	r, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		return
	}

	client, err := troutclient.CreateHTTPClient(s)
	if err != nil {
		fmt.Println("Error creating HTTP client:", err)
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(r)))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	req.Header.Set("Authorization", a)
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		fmt.Printf("%s user added successfully with password: %s\n", u, password)
	} else if resp.StatusCode == http.StatusForbidden {
		fmt.Println("Error response: Forbidden - must be 'admin' to add a user")
	} else if resp.StatusCode == http.StatusBadRequest {
		fmt.Printf("Username '%s' is already present.\n", data["username"])
	} else {
		fmt.Println("Error: User was not added successfully. Status code:", resp.StatusCode)
	}
}

func generatePassword() string {
	seed := time.Now().UnixNano()
	random := rand.New(rand.NewSource(seed))

	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()"
	password := make([]byte, 12)
	for i := range password {
		password[i] = charset[random.Intn(len(charset))]
	}
	return string(password)
}
