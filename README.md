#trout

## Overview
Trout is a CLI tool for interacting with the Gophish JSON API for phishing simulations

## Install

Clone from repo

```
$ go build
```
Tested with Go version 1.21.5+ 

## Usage
Trout has multiple options for the management of users, campaigns, landing pages, and email templates. All operations require the `-api-key` and `-gophish-server` flags.

```
$ ./trout -h


  O                         O
 oOo                       oOo
  o   `OoOo. .oOo. O   o    o
  O    o     O   o o   O    O
  o    O     o   O O   o    o
  `oO  o     `OoO' `OoO'o   `oO
			@h0useh3ad

Usage of ./trout:
  -add-emails
    	Add email templates
  -add-pages
    	Add landing page templates
  -add-user string
    	Username of the user to add
  -api-key string
    	Gophish API Key (required)
  -contains string
    	Add only template(s) that contain this string
  -delete-campaign string
    	Name of the campaign to delete (case-sensitive; check -list-campaigns)
  -delete-sending-profile string
    	Name of the sending profile to delete (case-sensitive; check -list-sending-profiles)
  -delete-user string
    	Username of the user to delete
  -gophish-server string
    	Gophish Server IP (required)
  -list-campaigns
    	List all campaigns
  -list-emails
    	List all email templates
  -list-pages
    	List all landing page templates
  -list-sending-profiles
    	List sending profiles
  -list-templates
    	List available templates
  -list-users
    	List all users
  -role string
    	Role of new user (user or admin)
  -socks string
    	SOCKS proxy address (ex: 127.0.0.1:1080)
  -version
    	Print tool version
```
### List users, landing pages, email templates, or sending profiles
`-list-users`, `-list-pages`, `-list-emails`, `-list-campaigns`, and `-list-sending-profiles` require the -api-key and -gophish-server flags. All user can list the landing pages, email templates, campaigns, or sending profiles for that user account. Only an 'admin' may -list-users.
```
./trout -api-key <key> -gophish-server <serverIP> -list-users
```
### List available templates
`-list-templates` checks the local directories `add/templates/emails` and `add/templates/pages` and lists the templates available (does not require -api-key or -gophish-server)
```
./trout -list-templates
```
##### Template file format
Additional email and landing page templates can be added by creating *.pagetemplate and *.emailtemplate files in their respective directories under `~/trout/add/templates/{emails,pages}`

The template files follow the necessary format for ingestion by the gophish API. Reference: https://docs.getgophish.com/api-documentation/templates
```
Example email template format:
{
  "name": "",
  "subject": "",
  "html": ""
}
```
### Add landing pages or email templates
`-add-pages` and `-add-emails` require the -api-key and -gophish-server flags. These options will populate target user account with each template located within the `add/templates/emails` and `add/templates/pages` directories by file extension. The -add-pages options will load each *.pagetemplate file, and -add-emails will load each *.emailtemplate file.
```
./trout -api-key <key> -gophish-server <serverIP> -add-emails
```
Adding the `-contains` flag will load only the email or landing page templates that contain a specific string in the name.
```
./trout -api-key <key> -gophish-server <serverIP> -add-emails -contains <string>
```
### Add a user
`-add-user` requires the -api-key -gophish-server -username and -role flags. The role of the user is either 'user' or 'admin'. A random 12-character password is generated for the user and provided in the display output. Only an 'admin' user may utilize the -add-user option
```
./trout -api-key <key> -gophish-server <serverIP> -username <username> -role user
```
### Delete user, campaign, or sending profile
`-delete-user`, `-delete-campaign`, and `-delete-sending-profile` require the -api-key and -gophish-server flags. Only an 'admin' user may utilize the -delete-user option. The -delete-campaign and -delete-sending-profile options can be used by any user account for its own workspace. 
```
./trout -api-key <key> -gophish-server <serverIP> -delete-user <username>
```
### SOCKS proxy
`-socks` can be added to your command to utilize a SOCKS proxy to access the Gophish API server.
```
./trout -api-key <key> -gophish-server <serverIP> -socks <socksAddress> -add-emails -contains <string>
```