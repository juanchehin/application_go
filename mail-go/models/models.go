package models

import "html/template"

// User is the user model
type User struct {
	To      string
	From    string
	Subject string
	Content template.HTML
}
