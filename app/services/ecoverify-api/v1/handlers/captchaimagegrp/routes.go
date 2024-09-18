package captchaimagegrp

import (
	"net/http"

	"github.com/SyedMohamedHyder/ecoverify/business/core/captchaimage"
	"github.com/SyedMohamedHyder/ecoverify/business/core/captchaimage/stores/captchaimagedb"
	"github.com/SyedMohamedHyder/ecoverify/business/web/v1/auth"
	"github.com/SyedMohamedHyder/ecoverify/business/web/v1/mid"
	"github.com/SyedMohamedHyder/ecoverify/foundation/logger"
	"github.com/SyedMohamedHyder/ecoverify/foundation/web"
	"github.com/jmoiron/sqlx"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Build string
	Log   *logger.Logger
	DB    *sqlx.DB
	Auth  *auth.Auth
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	const version = "v1"

	authen := mid.Authenticate(cfg.Auth)
	ruleAdmin := mid.Authorize(cfg.Auth, auth.RuleAdminOnly)

	captchaImageCore := captchaimage.NewCore(cfg.Log, captchaimagedb.NewStore(cfg.Log, cfg.DB))

	hdl := New(captchaImageCore, cfg.Auth)
	app.Handle(http.MethodPost, version, "/captchaimages", hdl.Create)
	app.Handle(http.MethodPost, version, "/captchaimagesauth", hdl.Create, authen, ruleAdmin)
}
