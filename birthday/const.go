package birthday

// BirthdateDefaultFormat is the borthdate format.
const BirthdateDefaultFormat = "02/01/2006"

// Provider is a provider.
type Provider struct {
	Type     string                 `json:"type"`
	Vendor   string                 `json:"vendor"`
	Settings map[string]interface{} `json:"settings"`
}

// Contact is a contact.
type Contact struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Birthdate string `json:"birthdate"`
}
