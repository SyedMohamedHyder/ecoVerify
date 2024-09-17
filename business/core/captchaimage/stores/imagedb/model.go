package imagedb

import (
	"fmt"
	"time"

	"github.com/SyedMohamedHyder/ecoverify/business/core/captchaimage"
	"github.com/google/uuid"
)

// dbCaptchaImage represent the structure we need for moving data
// between the app and the database.
type dbCaptchaImage struct {
	ID          uuid.UUID `db:"image_id"`
	URL         string    `db:"url"`
	Category    string    `db:"image_category"`
	DateCreated time.Time `db:"date_created"`
	DateUpdated time.Time `db:"date_updated"`
}

// toDBCaptchaImage transforms a core image into a database image.
func toDBCaptchaImage(img captchaimage.CaptchaImage) dbCaptchaImage {
	return dbCaptchaImage{
		ID:          img.ID,
		URL:         img.URL,
		Category:    img.Category.Name(),
		DateCreated: img.DateCreated.UTC(),
		DateUpdated: img.DateUpdated.UTC(),
	}
}

func toCoreCaptchaImage(dbImg dbCaptchaImage) (captchaimage.CaptchaImage, error) {
	category, err := captchaimage.ParseCaptchaImageCategory(dbImg.Category)
	if err != nil {
		return captchaimage.CaptchaImage{}, fmt.Errorf("parse captcha image category: %w", err)
	}

	img := captchaimage.CaptchaImage{
		ID:          dbImg.ID,
		URL:         dbImg.URL,
		Category:    category,
		DateCreated: dbImg.DateCreated.In(time.Local),
		DateUpdated: dbImg.DateUpdated.In(time.Local),
	}

	return img, nil
}

func toCoreUserSlice(dbImgs []dbCaptchaImage) ([]captchaimage.CaptchaImage, error) {
	imgs := make([]captchaimage.CaptchaImage, len(dbImgs))

	for i, dbImg := range dbImgs {
		var err error
		imgs[i], err = toCoreCaptchaImage(dbImg)
		if err != nil {
			return nil, err
		}
	}

	return imgs, nil
}
