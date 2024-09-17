package image

import (
	"fmt"
	"time"

	"github.com/SyedMohamedHyder/ecoverify/foundation/validate"
	"github.com/google/uuid"
)

// QueryFilter holds the available fields a query can be filtered on.
type QueryFilter struct {
	ID               *uuid.UUID `validate:"omitempty"`
	URL              *string    `validate:"omitempty,min=3"`
	Type             *ImageType `validate:"omitempty"`
	StartCreatedDate *time.Time `validate:"omitempty"`
	EndCreatedDate   *time.Time `validate:"omitempty"`
}

// Validate checks the data in the model is considered clean.
func (qf *QueryFilter) Validate() error {
	if err := validate.Check(qf); err != nil {
		return fmt.Errorf("validate: %w", err)
	}
	return nil
}

// WithImageID sets the ID field of the QueryFilter value.
func (qf *QueryFilter) WithImageID(imageID uuid.UUID) {
	qf.ID = &imageID
}

// WithURL sets the URL field of the QueryFilter value.
func (qf *QueryFilter) WithURL(url string) {
	qf.URL = &url
}

// WithImageType sets the ImageType field of the QueryFilter value.
func (qf *QueryFilter) WithImageType(imgType ImageType) {
	qf.Type = &imgType
}

// WithStartDateCreated sets the DateCreated field of the QueryFilter value.
func (qf *QueryFilter) WithStartDateCreated(startDate time.Time) {
	d := startDate.UTC()
	qf.StartCreatedDate = &d
}

// WithEndCreatedDate sets the DateCreated field of the QueryFilter value.
func (qf *QueryFilter) WithEndCreatedDate(endDate time.Time) {
	d := endDate.UTC()
	qf.EndCreatedDate = &d
}
