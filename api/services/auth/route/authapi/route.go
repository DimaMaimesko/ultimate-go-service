package authapi

import (
	"github.com/DimaMaimesko/ultimate-go-service/api/services/api/mid"
	"github.com/DimaMaimesko/ultimate-go-service/business/api/auth"
	"github.com/DimaMaimesko/ultimate-go-service/foundation/web"
)

// Routes adds specific routes for this group.
func Routes(app *web.App, a *auth.Auth) {
	authen := mid.AuthenticateLocal(a)

	api := newAPI(a)
	app.HandleFunc("GET /auth/token/{kid}", api.token, authen)
	app.HandleFunc("GET /auth/authenticate", api.authenticate, authen)
	app.HandleFunc("POST /auth/authorize", api.authorize)
}
