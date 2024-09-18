package captchaimagedb

import (
	"context"
	"errors"
	"fmt"

	"github.com/SyedMohamedHyder/ecoverify/business/core/captchaimage"
	db "github.com/SyedMohamedHyder/ecoverify/business/data/dbsql/pgx"
	"github.com/SyedMohamedHyder/ecoverify/foundation/logger"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

// Store manages the set of APIs for captcha image database access.
type Store struct {
	log *logger.Logger
	db  *sqlx.DB
}

// NewStore constructs the api for data access.
func NewStore(log *logger.Logger, db *sqlx.DB) *Store {
	return &Store{
		log: log,
		db:  db,
	}
}

// Create inserts a new captcha image into the database.
func (s *Store) Create(ctx context.Context, img captchaimage.CaptchaImage) error {
	const q = `
	INSERT INTO captcha_images
		(image_id, url, image_category, date_created, date_updated)
	VALUES
		(:image_id, :url, :image_category, :date_created, :date_updated)`

	if err := db.NamedExecContext(ctx, s.log, s.db, q, toDBCaptchaImage(img)); err != nil {
		if errors.Is(err, db.ErrDBDuplicatedEntry) {
			return fmt.Errorf("namedexeccontext: %w", captchaimage.ErrUniqueURL)
		}
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

// QueryByID gets the specified captcha image from the database.
func (s *Store) QueryByID(ctx context.Context, imgID uuid.UUID) (captchaimage.CaptchaImage, error) {
	data := struct {
		ID string `db:"image_id"`
	}{
		ID: imgID.String(),
	}

	const q = `
	SELECT
        image_id, url, image_category, date_created, date_updated
	FROM
		captcha_images
	WHERE 
		image_id = :image_id`

	var dbImg dbCaptchaImage
	if err := db.NamedQueryStruct(ctx, s.log, s.db, q, data, &dbImg); err != nil {
		if errors.Is(err, db.ErrDBNotFound) {
			return captchaimage.CaptchaImage{}, fmt.Errorf("namedquerystruct: %w", captchaimage.ErrNotFound)
		}
		return captchaimage.CaptchaImage{}, fmt.Errorf("namedquerystruct: %w", err)
	}

	img, err := toCoreCaptchaImage(dbImg)
	if err != nil {
		return captchaimage.CaptchaImage{}, err
	}

	return img, nil
}

// QueryByURL gets the specified captcha image from the database by URL.
func (s *Store) QueryByURL(ctx context.Context, url string) (captchaimage.CaptchaImage, error) {
	data := struct {
		URL string `db:"url"`
	}{
		URL: url,
	}

	const q = `
	SELECT
        image_id, url, image_category, date_created, date_updated
	FROM
		captcha_images
	WHERE
		url = :url`

	var dbImg dbCaptchaImage
	if err := db.NamedQueryStruct(ctx, s.log, s.db, q, data, &dbImg); err != nil {
		if errors.Is(err, db.ErrDBNotFound) {
			return captchaimage.CaptchaImage{}, fmt.Errorf("namedquerystruct: %w", captchaimage.ErrNotFound)
		}
		return captchaimage.CaptchaImage{}, fmt.Errorf("namedquerystruct: %w", err)
	}

	img, err := toCoreCaptchaImage(dbImg)
	if err != nil {
		return captchaimage.CaptchaImage{}, err
	}

	return img, nil
}
