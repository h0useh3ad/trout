package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	troutadd "trout/add"
	troutdelete "trout/delete"
	troutlist "trout/list"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
)

const version = "1.0"

var showVersion bool
var apiKey string
var phishServer string
var contains string
var socksAddr string

func main() {
	figure.NewColorFigure("trout", "pebbles", "green", true).Print()
	color.Green("\t\t\t@h0useh3ad\n\n")

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-interrupt
		fmt.Println("\nExiting...")
		os.Exit(0)
	}()

	flag.BoolVar(&showVersion, "version", false, "Print tool version")
	flag.StringVar(&apiKey, "api-key", "", "Gophish API Key (required)")
	flag.StringVar(&phishServer, "gophish-server", "", "Gophish Server IP (required)")
	listUsers := flag.Bool("list-users", false, "List all users")
	listEmails := flag.Bool("list-emails", false, "List all email templates")
	listPages := flag.Bool("list-pages", false, "List all landing page templates")
	listCampaigns := flag.Bool("list-campaigns", false, "List all campaigns")
	listTemplates := flag.Bool("list-templates", false, "List available templates")
	listSmtp := flag.Bool("list-sending-profiles", false, "List sending profiles")
	addUser := flag.String("add-user", "", "Username of the user to add")
	addEmails := flag.Bool("add-emails", false, "Add email templates")
	addPages := flag.Bool("add-pages", false, "Add landing page templates")
	flag.StringVar(&contains, "contains", "", "Add only template(s) that contain this string")
	deleteUsername := flag.String("delete-user", "", "Username of the user to delete")
	deleteCampaign := flag.String("delete-campaign", "", "Name of the campaign to delete (case-sensitive; check -list-campaigns)")
	deleteSmtp := flag.String("delete-sending-profile", "", "Name of the sending profile to delete (case-sensitive; check -list-sending-profiles)")
	role := flag.String("role", "", "Role of new user (user or admin)")
	flag.StringVar(&socksAddr, "socks", "", "SOCKS proxy address (ex: 127.0.0.1:1080)")
	flag.Parse()

	if showVersion {
		fmt.Println("Version:", version)
		os.Exit(0)
	}

	if *listTemplates {
		troutlist.GetTemplates("add/templates/emails", "Email", ".emailtemplate")
		troutlist.GetTemplates("add/templates/pages", "Landing page", ".pagetemplate")
	} else {

		if apiKey == "" || phishServer == "" {
			fmt.Println("API Key & GoPhish Server IP is required")
			os.Exit(1)
		}

		if *listUsers {
			troutlist.GetUsers(apiKey, phishServer, socksAddr)
		}

		if *listEmails {
			troutlist.GetEmails(apiKey, phishServer, socksAddr)
		}

		if *listPages {
			troutlist.GetPages(apiKey, phishServer, socksAddr)
		}

		if *listCampaigns {
			troutlist.GetCampaigns(apiKey, phishServer, socksAddr)
		}

		if *listSmtp {
			troutlist.GetSmtp(apiKey, phishServer, socksAddr)
		}

		if *addUser != "" {
			if *role == "" {
				fmt.Println("Role is required for adding a new user")
				os.Exit(1)
			}
			troutadd.AddUser(apiKey, phishServer, *addUser, *role, socksAddr)
		}

		if *addEmails {
			if contains == "" {
				troutadd.AddEmails(apiKey, phishServer, socksAddr)
			} else {
				troutadd.AddEmailContains(apiKey, phishServer, contains, socksAddr)
			}
		}

		if *addPages {
			if contains == "" {
				troutadd.AddPages(apiKey, phishServer, socksAddr)
			} else {
				troutadd.AddPageContains(apiKey, phishServer, contains, socksAddr)
			}
		}

		if *deleteUsername != "" {
			troutdelete.DeleteUser(apiKey, phishServer, *deleteUsername, socksAddr)
		}

		if *deleteCampaign != "" {
			troutdelete.DeleteCampaign(apiKey, phishServer, *deleteCampaign, socksAddr)
		}

		if *deleteSmtp != "" {
			troutdelete.DeleteSmtp(apiKey, phishServer, *deleteSmtp, socksAddr)
		}
	}
}
