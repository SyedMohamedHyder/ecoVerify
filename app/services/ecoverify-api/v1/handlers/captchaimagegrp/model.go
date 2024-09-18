package captchaimagegrp

import (
	"fmt"
	"time"

	"github.com/SyedMohamedHyder/ecoverify/business/core/captchaimage"
	"github.com/SyedMohamedHyder/ecoverify/foundation/validate"
)

// AppCaptchaImage represents information about an individual captcha image.
type AppCaptchaImage struct {
	ID          string `json:"id"`
	URL         string `json:"url"`
	Category    string `json:"category"`
	DateCreated string `json:"dateCreated"`
	DateUpdated string `json:"dateUpdated"`
}

func toAppCaptchaImage(img captchaimage.CaptchaImage) AppCaptchaImage {
	return AppCaptchaImage{
		ID:          img.ID.String(),
		URL:         img.URL,
		Category:    img.Category.Name(),
		DateCreated: img.DateCreated.Format(time.RFC3339),
		DateUpdated: img.DateUpdated.Format(time.RFC3339),
	}
}

func toAppCaptchaImages(imgs []captchaimage.CaptchaImage) []AppCaptchaImage {
	items := make([]AppCaptchaImage, len(imgs))
	for i, img := range imgs {
		items[i] = toAppCaptchaImage(img)
	}

	return items
}

// =============================================================================

// AppNewCaptchaImage contains information needed to create a new captcha image.
type AppNewCaptchaImage struct {
	URL      string `json:"url" validate:"required"`
	Category string `json:"category" validate:"required"`
}

func toCoreNewCaptchaImage(app AppNewCaptchaImage) (captchaimage.NewCaptchaImage, error) {
	category, err := captchaimage.ParseCaptchaImageCategory(app.Category)
	if err != nil {
		return captchaimage.NewCaptchaImage{}, fmt.Errorf("parsing captcha image category: %w", err)
	}

	img := captchaimage.NewCaptchaImage{
		URL:      app.URL,
		Category: category,
	}

	return img, nil
}

// Validate checks the data in the model is considered clean.
func (app AppNewCaptchaImage) Validate() error {
	if err := validate.Check(app); err != nil {
		return err
	}

	return nil
}
