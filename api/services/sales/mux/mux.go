package mux

import (
	"os"

	"github.com/DimaMaimesko/ultimate-go-service/api/services/api/mid"
	"github.com/DimaMaimesko/ultimate-go-service/api/services/sales/route/sys/checkapi"
	"github.com/DimaMaimesko/ultimate-go-service/app/api/authclient"
	"github.com/DimaMaimesko/ultimate-go-service/foundation/logger"
	"github.com/DimaMaimesko/ultimate-go-service/foundation/web"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(log *logger.Logger, authClient *authclient.Client, shutdown chan os.Signal) *web.App {
	app := web.NewApp(shutdown, mid.Logger(log), mid.Errors(log), mid.Metrics(), mid.Panics())

	checkapi.Routes(app, log, authClient)

	return app
}
