// Package mux provides support to bind domain level routes
// to the application mux.
package mux

import (
	"os"

	"github.com/DimaMaimesko/ultimate-go-service/api/services/api/mid"
	"github.com/DimaMaimesko/ultimate-go-service/api/services/auth/route/authapi"
	"github.com/DimaMaimesko/ultimate-go-service/api/services/auth/route/checkapi"
	"github.com/DimaMaimesko/ultimate-go-service/business/api/auth"
	"github.com/DimaMaimesko/ultimate-go-service/foundation/logger"
	"github.com/DimaMaimesko/ultimate-go-service/foundation/web"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(log *logger.Logger, auth *auth.Auth, shutdown chan os.Signal) *web.App {
	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics())

	checkapi.Routes(app, auth)
	authapi.Routes(app, auth)

	return app
}
