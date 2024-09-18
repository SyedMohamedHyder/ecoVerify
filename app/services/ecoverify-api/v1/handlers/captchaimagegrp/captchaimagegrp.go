package captchaimagegrp

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/SyedMohamedHyder/ecoverify/business/core/captchaimage"
	"github.com/SyedMohamedHyder/ecoverify/business/web/v1/auth"
	"github.com/SyedMohamedHyder/ecoverify/business/web/v1/response"
	"github.com/SyedMohamedHyder/ecoverify/foundation/web"
)

// Handlers manages the set of captchaimage endpoints.
type Handlers struct {
	captchaImage *captchaimage.Core
	auth         *auth.Auth
}

// New constructs a handlers for route access.
func New(img *captchaimage.Core, auth *auth.Auth) *Handlers {
	return &Handlers{
		captchaImage: img,
		auth:         auth,
	}
}

// Create adds a new user to the system.
func (h *Handlers) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var app AppNewCaptchaImage
	if err := web.Decode(r, &app); err != nil {
		return response.NewError(err, http.StatusBadRequest)
	}

	nci, err := toCoreNewCaptchaImage(app)
	if err != nil {
		return response.NewError(err, http.StatusBadRequest)
	}

	img, err := h.captchaImage.Create(ctx, nci)
	if err != nil {
		if errors.Is(err, captchaimage.ErrUniqueURL) {
			return response.NewError(err, http.StatusConflict)
		}
		return fmt.Errorf("create: captchaimage[%+v]: %w", img, err)
	}

	return web.Respond(ctx, w, toAppCaptchaImage(img), http.StatusCreated)
}
