package youremailapigosdk

type SendEmailData struct {
	Template    string            `json:"template"`
	SmtpAccount string            `json:"smtp_account"`
	Subject     string            `json:"subject"`
	ReplyTo     map[string]string `json:"reply_to"`
	Variables   map[string]string `json:"variables"`
	To          string            `json:"to"`
}
