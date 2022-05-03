package config

import (
	"mail-go/models"
)

// AppConfig holds the application config
type AppConfig struct {
	MailChan chan models.MailData
}
