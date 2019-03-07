package providers

import (
	"time"
)

// Provider interface
type Provider interface {
	SendNotification(string, string, time.Time) error
}
