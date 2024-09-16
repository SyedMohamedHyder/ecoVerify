package handlers

import (
	"github.com/SyedMohamedHyder/ecoVerify/app/services/eco-verify-api/v1/handlers/hackgrp"
	v1 "github.com/SyedMohamedHyder/ecoVerify/business/web/v1"
	"github.com/dimfeld/httptreemux/v5"
)

type Routes struct{}

// Add implements the RouterAdder interface.
func (Routes) Add(mux *httptreemux.ContextMux, cfg v1.APIMuxConfig) {
	hackgrp.Routes(mux)
}
