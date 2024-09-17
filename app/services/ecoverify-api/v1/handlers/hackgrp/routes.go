package hackgrp

import (
	"net/http"

	"github.com/SyedMohamedHyder/ecoverify/business/web/v1/auth"
	"github.com/SyedMohamedHyder/ecoverify/business/web/v1/mid"
	"github.com/SyedMohamedHyder/ecoverify/foundation/web"
)

// Config contains all the mandatory systems required by handlers.
type Config struct {
	Auth *auth.Auth
}

// Routes adds specific routes for this group.
func Routes(app *web.App, cfg Config) {
	authen := mid.Authenticate(cfg.Auth)
	ruleAdmin := mid.Authorize(cfg.Auth, auth.RuleAdminOnly)

	app.Handle(http.MethodGet, "/hack", Hack)
	app.Handle(http.MethodGet, "/hackauth", Hack, authen, ruleAdmin)
}
