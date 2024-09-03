package troutstructs

type Forbidden struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

type User struct {
	ID             int         `json:"id"`
	Username       string      `json:"username"`
	Password       string      `json:"password"`
	ApiKey         string      `json:"api_key"`
	Role           RoleDetails `json:"role"`
	PassChangeReqd bool        `json:"password_change_reqd"`
	AccountLocked  bool        `json:"account_locked"`
	LastLogin      string      `json:"last_login"`
}

type RoleDetails struct {
	Slug        string `json:"slug"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type EmailTemplate struct {
	ID             int          `json:"id"`
	Name           string       `json:"name"`
	EnvelopeSender string       `json:"envelope_sender"`
	Subject        string       `json:"subject"`
	Text           string       `json:"text"`
	Html           string       `json:"html"`
	ModifiedDate   string       `json:"modified_date"`
	Attachments    []Attachment `json:"attachments"`
}

type Attachment struct {
	Content string `json:"content"`
	Type    string `json:"type"`
	Name    string `json:"name"`
}

type PageTemplate struct {
	ID                 int    `json:"id"`
	Name               string `json:"name"`
	Html               string `json:"html"`
	CaptureCredentials bool   `json:"capture_credentials"`
	CapturePassswords  bool   `json:"capture_passwords"`
	ModifiedDate       string `json:"modified_date"`
	RedirectURL        string `json:"redirect_url"`
}

type Campaigns struct {
	ID            int               `json:"id"`
	Name          string            `json:"name"`
	CreatedDate   string            `json:"created_date"`
	LaunchDate    string            `json:"launch_date"`
	SendByDate    string            `json:"send_by_date"`
	CompletedDate string            `json:"completed_date"`
	Template      EmailTemplate     `json:"template"`
	Page          PageTemplate      `json:"page"`
	Status        string            `json:"status"`
	Results       []CampaignResults `json:"results"`
	Groups        string            `json:"groups"`
	Timeline      []CampaignEvents  `json:"timeline"`
	Smtp          SMTP              `json:"smtp"`
	Url           string            `json:"url"`
}

type CampaignEvents struct {
	Email   string `json:"email"`
	Time    string `json:"time"`
	Message string `json:"message"`
	Details string `json:"details"`
}

type CampaignResults struct {
	ID        string  `json:"id"`
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Position  string  `json:"position"`
	Status    string  `json:"status"`
	IP        string  `json:"ip"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	SendDate  string  `json:"send_date"`
	Reported  bool    `json:"reported"`
}

type SMTP struct {
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	Username         string   `json:"username"`
	Password         string   `json:"password"`
	Host             string   `json:"host"`
	InterfaceType    string   `json:"interface_type"`
	FromAddress      string   `json:"from_address"`
	IgnoreCertErrors bool     `json:"ignore_cert_errors"`
	ModifiedDate     string   `json:"modified_date"`
	Headers          []Header `json:"headers,omitempty"`
}

type Header struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DeleteResponse struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}
