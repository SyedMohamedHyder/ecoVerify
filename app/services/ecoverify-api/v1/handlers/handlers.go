package handlers

import (
	"github.com/SyedMohamedHyder/ecoverify/app/services/ecoverify-api/v1/handlers/captchaimagegrp"
	"github.com/SyedMohamedHyder/ecoverify/app/services/ecoverify-api/v1/handlers/checkgrp"
	"github.com/SyedMohamedHyder/ecoverify/app/services/ecoverify-api/v1/handlers/hackgrp"
	v1 "github.com/SyedMohamedHyder/ecoverify/business/web/v1"
	"github.com/SyedMohamedHyder/ecoverify/foundation/web"
)

type Routes struct{}

// Add implements the RouterAdder interface to add all routes.
func (Routes) Add(app *web.App, apiCfg v1.APIMuxConfig) {
	hackgrp.Routes(app, hackgrp.Config{
		Auth: apiCfg.Auth,
	})

	checkgrp.Routes(app, checkgrp.Config{
		Build: apiCfg.Build,
		Log:   apiCfg.Log,
	})

	captchaimagegrp.Routes(
		app,
		captchaimagegrp.Config{
			Build: apiCfg.Build,
			Log:   apiCfg.Log,
			DB:    apiCfg.DB,
			Auth:  apiCfg.Auth,
		},
	)
}
