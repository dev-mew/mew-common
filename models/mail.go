package models

type MailData struct {
	To      string
	CC      []string
	BCC     []string
	Subject string
	Content string
}
