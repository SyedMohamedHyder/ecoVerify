-- Version: 1.01
-- Description: Create table images
CREATE TABLE captcha_images (
	image_id           UUID    		     NOT NULL,
	url                TEXT    UNIQUE    NOT NULL,
	image_category     TEXT        		 NOT NULL,
	date_created      TIMESTAMP    		 NOT NULL,
	date_updated      TIMESTAMP    		 NOT NULL,

	PRIMARY KEY (image_id)
);
