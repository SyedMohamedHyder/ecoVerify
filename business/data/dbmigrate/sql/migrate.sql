-- Version: 1.01
-- Description: Create table images
CREATE TABLE images (
	image_id       UUID        NOT NULL,
	url            TEXT        NOT NULL,
	image_type     TEXT        NOT NULL,
	date_created  TIMESTAMP    NOT NULL,
	date_updated  TIMESTAMP    NOT NULL,

	PRIMARY KEY (image_id)
);
