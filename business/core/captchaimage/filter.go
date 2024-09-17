package captchaimage

import (
	"fmt"
	"time"

	"github.com/SyedMohamedHyder/ecoverify/foundation/validate"
	"github.com/google/uuid"
)

// QueryFilter holds the available fields a query can be filtered on.
type QueryFilter struct {
	ID               *uuid.UUID            `validate:"omitempty"`
	URL              *string               `validate:"omitempty,min=3"`
	Category         *CaptchaImageCategory `validate:"omitempty"`
	StartCreatedDate *time.Time            `validate:"omitempty"`
	EndCreatedDate   *time.Time            `validate:"omitempty"`
}

// Validate checks the data in the model is considered clean.
func (qf *QueryFilter) Validate() error {
	if err := validate.Check(qf); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}

// WithCaptchaImageID sets the ID field of the QueryFilter value.
func (qf *QueryFilter) WithCaptchaImageID(imageID uuid.UUID) {
	qf.ID = &imageID
}

// WithCaptchaImageURL sets the URL field of the QueryFilter value.
func (qf *QueryFilter) WithCaptchaImageURL(url string) {
	qf.URL = &url
}

// WithCaptchaImageCategory sets the Category field of the QueryFilter value.
func (qf *QueryFilter) WithCaptchaImageCategory(imgCategory CaptchaImageCategory) {
	qf.Category = &imgCategory
}

// WithCaptchaImageStartDateCreated sets the DateCreated field of the QueryFilter value.
func (qf *QueryFilter) WithCaptchaImageStartDateCreated(startDate time.Time) {
	d := startDate.UTC()
	qf.StartCreatedDate = &d
}

// WithCaptchaImageEndCreatedDate sets the DateCreated field of the QueryFilter value.
func (qf *QueryFilter) WithCaptchaImageEndCreatedDate(endDate time.Time) {
	d := endDate.UTC()
	qf.EndCreatedDate = &d
}
