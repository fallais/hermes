package providers

// Provider interface
type Provider interface {
	SendNotification(string, string, int) error
	Type() string
	Vendor() string
}
