// captcha.go - Captcha handling strategies
package juice

import (
	"image"
	"io"
)

// CaptchaMode defines how captchas are handled
type CaptchaMode int

const (
	// CaptchaModeManual returns the captcha for manual solving
	CaptchaModeManual CaptchaMode = iota

	// CaptchaModeOCR attempts to solve captcha with built-in OCR
	CaptchaModeOCR

	// CaptchaModeService uses a third-party service who knows?
	CaptchaModeService
)

// CaptchaResult represents a captcha challenge
type CaptchaResult struct {
	Image       image.Image // For manual mode
	ImageBytes  []byte      // Raw image bytes
	Solution    string      // For automated modes
	NeedsManual bool        // Indicates if user input is needed
}

// CaptchaSolver interface for different solving strategies
type CaptchaSolver interface {
	Solve(captchaImg io.Reader) (*CaptchaResult, error)
}

// GetCaptcha retrieves a captcha and processes it according to mode
func (c *Client) GetCaptcha() (*CaptchaResult, error) {
	return nil, nil
}

// manualCaptchaSolver implements CaptchaSolver for manual solving
type manualCaptchaSolver struct{}

func (s *manualCaptchaSolver) Solve(captchaImg io.Reader) (*CaptchaResult, error) {
	imgBytes, err := io.ReadAll(captchaImg)
	if err != nil {
		return nil, err
	}

	return &CaptchaResult{
		ImageBytes:  imgBytes,
		NeedsManual: true,
	}, nil
}
