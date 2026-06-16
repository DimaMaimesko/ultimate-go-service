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
	"github.com/jmoiron/sqlx"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(build string, log *logger.Logger, db *sqlx.DB, auth *auth.Auth, shutdown chan os.Signal) *web.App {
	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics())

	checkapi.Routes(build, app, log, db)
	authapi.Routes(app, auth)

	return app
}
