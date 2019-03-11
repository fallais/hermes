package providers

import (
	"gobirthday/models"
)

// Provider interface
type Provider interface {
	SendNotification(*models.Contact) error
	Type() string
	Vendor() string
}
