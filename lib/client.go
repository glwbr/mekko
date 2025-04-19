// client.go - Core API client with HTTP handling
package juice

import (
	"net/http"
	"time"
)

// Client handles all interactions with the SEFAZ API
type Client struct {
	httpClient *http.Client
	cookies    []*http.Cookie
	userAgent  string
	baseURL    string
	logger     Logger
	options    *ClientOptions
}

// ClientOptions provides configuration for the SEFAZ client
type ClientOptions struct {
	Timeout       time.Duration
	UserAgent     string
	CaptchaMode   CaptchaMode
	Logger        Logger
	DebugMode     bool
	RetryAttempts int
}

// New creates a configured SEFAZ Client
func New(options *ClientOptions) *Client {
	if options == nil {
		options = &ClientOptions{}
	}

	if options.Timeout == 0 {
		options.Timeout = 10 * time.Second
	}

	if options.UserAgent == "" {
		options.UserAgent = DefaultUserAgent
	}

	return &Client{
		httpClient: &http.Client{Timeout: options.Timeout},
		userAgent:  options.UserAgent,
		baseURL:    SEFAZEndpoint,
		options:    options,
		logger:     options.Logger,
	}
}

// GetInvoice is the entrypoint for retrieving invoice data
func (c *Client) GetInvoice(accessKey string) (*Invoice, error) {
	// Implementation steps:
	// 1. Get home page to establish session
	// 2. Get captcha if needed
	// 3. Process captcha according to CaptchaMode
	// 4. Submit form with access key and captcha
	// 5. Parse response to Invoice model
	// 6. Return structured data

	return nil, nil
}
