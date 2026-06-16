package mux

import (
	"os"

	"github.com/DimaMaimesko/ultimate-go-service/api/services/api/mid"
	"github.com/DimaMaimesko/ultimate-go-service/api/services/sales/route/sys/checkapi"
	"github.com/DimaMaimesko/ultimate-go-service/app/api/authclient"
	"github.com/DimaMaimesko/ultimate-go-service/foundation/logger"
	"github.com/DimaMaimesko/ultimate-go-service/foundation/web"
	"github.com/jmoiron/sqlx"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(build string, log *logger.Logger, db *sqlx.DB, authClient *authclient.Client, shutdown chan os.Signal) *web.App {
	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics())

	checkapi.Routes(build, app, log, db, authClient)

	return app
}
