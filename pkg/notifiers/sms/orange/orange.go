package orange

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"hermes/pkg/notifiers"
	"hermes/pkg/notifiers/sms"

	"github.com/sirupsen/logrus"
)

//------------------------------------------------------------------------------
// Structure
//------------------------------------------------------------------------------

// Type is the type of the provider.
const Type = "SMS"

// Vendor is the vendor of the provider.
const Vendor = "Orange"

type orange struct {
	client *http.Client
	url    string
	user   string
	pass   string
}

//------------------------------------------------------------------------------
// Factory
//------------------------------------------------------------------------------

// NewNotifier returns a new notifier for Free.
func NewNotifier(settings map[string]interface{}) notifiers.Notifier {
	// Set HTTP transport
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	// Set timeout
	timeout := time.Duration(12 * time.Second)

	// Set HTTP Client
	client := &http.Client{
		Transport: tr,
		Timeout:   timeout,
	}

	// Initial values
	user := ""
	pass := ""

	// Process the values
	for key, value := range settings {
		switch key {
		case "user":
			user = value.(string)
		case "pass":
			pass = value.(string)
		default:
			logrus.WithFields(logrus.Fields{
				"key":   key,
				"value": value,
			}).Infoln("Wrong setting for Free")
		}
	}

	return &orange{
		client: client,
		url:    "https://api.orange.com",
		user:   user,
		pass:   pass,
	}
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// Notify sends a notification.
func (n *orange) Notify(msg string) error {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(n.url)
	if err != nil {
		return fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/oauth/v2/token"
	parameters := url.Values{}
	parameters.Add("grant_type", "client_credentials")
	parameters.Add("pass", n.pass)
	parameters.Add("msg", msg)
	reqURL.RawQuery = parameters.Encode()

	// Create the request
	req, err := http.NewRequest("GET", reqURL.String(), nil)
	if err != nil {
		return fmt.Errorf("Error while creating the request : %s", err)
	}

	// Do the request
	resp, err := n.client.Do(req)
	if err != nil {
		return fmt.Errorf("Error while doing the request : %s", err)
	}
	defer resp.Body.Close()

	// Check the status code
	if resp.StatusCode != 200 {
		return fmt.Errorf("Error while sending the SMS. Status code is %d", resp.StatusCode)
	}

	return nil
}

// Name returns the formatted name.
func (n *orange) Name() string {
	return fmt.Sprintf("%s-%s", sms.Type, Vendor)
}
