package models

import "html/template"

//
type MailData struct {
	To      string
	From    string
	Subject string
	Content template.HTML
}
