package image

import (
	"context"
	"fmt"
	"time"

	"github.com/SyedMohamedHyder/ecoverify/foundation/logger"
	"github.com/google/uuid"
)

// Storer interface declares the behavior this package needs to persists and
// retrieve data.
type Storer interface {
	Create(ctx context.Context, img Image) error
}

// =============================================================================

// Core manages the set of APIs for user access.
type Core struct {
	storer Storer
	log    *logger.Logger
}

// NewCore constructs a core for user api access.
func NewCore(log *logger.Logger, storer Storer) *Core {
	return &Core{
		storer: storer,
		log:    log,
	}
}

// Create adds a new user to the system.
func (c *Core) Create(ctx context.Context, ni NewImage) (Image, error) {
	now := time.Now()

	img := Image{
		ID:          uuid.New(),
		URL:         ni.URL,
		Type:        ni.Type,
		DateCreated: now,
		DateUpdated: now,
	}

	if err := c.storer.Create(ctx, img); err != nil {
		return Image{}, fmt.Errorf("create: %w", err)
	}

	return img, nil
}
