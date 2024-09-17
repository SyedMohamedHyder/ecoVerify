package captchaimage

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/SyedMohamedHyder/ecoverify/foundation/logger"
	"github.com/google/uuid"
)

// Set of error variables for CRUD operations.
var (
	ErrNotFound              = errors.New("image not found")
	ErrUniqueURL             = errors.New("url is not unique")
	ErrAuthenticationFailure = errors.New("authentication failed")
)

// Storer interface declares the behavior this package needs to persists and
// retrieve data.
type Storer interface {
	Create(ctx context.Context, img CaptchaImage) error
}

// =============================================================================

// Core manages the set of APIs for captcha image access.
type Core struct {
	storer Storer
	log    *logger.Logger
}

// NewCore constructs a core for captcha image api access.
func NewCore(log *logger.Logger, storer Storer) *Core {
	return &Core{
		storer: storer,
		log:    log,
	}
}

// Create adds a new captcha image to the system.
func (c *Core) Create(ctx context.Context, ni NewCaptchaImage) (CaptchaImage, error) {
	now := time.Now()

	img := CaptchaImage{
		ID:          uuid.New(),
		URL:         ni.URL,
		Category:    ni.Category,
		DateCreated: now,
		DateUpdated: now,
	}

	if err := c.storer.Create(ctx, img); err != nil {
		return CaptchaImage{}, fmt.Errorf("create: %w", err)
	}

	return img, nil
}
