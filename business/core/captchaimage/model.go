package captchaimage

import (
	"time"

	"github.com/google/uuid"
)

// CaptchaImage represents information about an individual image used in CAPTCHA.
type CaptchaImage struct {
	ID          uuid.UUID            // Unique identifier for each image
	URL         string               // URL pointing to the image location
	Category    CaptchaImageCategory // Category of image (eco-friendly or non-eco-friendly)
	DateCreated time.Time            // Timestamp when the image was created
	DateUpdated time.Time            // Timestamp when the image was last updated
}

// NewCaptchaImage contains information needed to create a new captcha image in the database.
type NewCaptchaImage struct {
	URL      string               // URL for the new image
	Category CaptchaImageCategory // Category of the image (eco-friendly or non-eco-friendly)
}

// UpdateCaptchaImage contains information needed to update a captcha image.
type UpdateCaptchaImage struct {
	URL      *string               // New URL for the image (optional)
	Category *CaptchaImageCategory // New Category for the image (optional)
}
