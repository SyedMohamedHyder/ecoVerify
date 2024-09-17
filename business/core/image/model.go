package image

import (
	"time"

	"github.com/google/uuid"
)

// ImageType defines types of images used in the CAPTCHA.
type ImageType string

const (
	ImageTypeEcoFriendly    ImageType = "ecofriendly"    // Represents eco-friendly images
	ImageTypeNonEcoFriendly ImageType = "nonecofriendly" // Represents non-eco-friendly images
)

// Image represents information about an individual image used in CAPTCHA.
type Image struct {
	ID          uuid.UUID // Unique identifier for each image
	URL         string    // URL pointing to the image location
	Type        ImageType // Type of image (eco-friendly or non-eco-friendly)
	DateCreated time.Time // Timestamp when the image was created
	DateUpdated time.Time // Timestamp when the image was last updated
}

// NewImage contains information needed to create a new image in the database.
type NewImage struct {
	URL  string    // URL for the new image
	Type ImageType // Type of the image (eco-friendly or non-eco-friendly)
}

// UpdateImage contains information needed to update an image.
type UpdateImage struct {
	URL  *string    // New URL for the image (optional)
	Type *ImageType // New type for the image (optional)
}
