package checkapi

import (
	"github.com/DimaMaimesko/ultimate-go-service/api/services/api/mid"
	"github.com/DimaMaimesko/ultimate-go-service/app/api/authclient"
	"github.com/DimaMaimesko/ultimate-go-service/business/api/auth"
	"github.com/DimaMaimesko/ultimate-go-service/foundation/logger"
	"github.com/DimaMaimesko/ultimate-go-service/foundation/web"
	"github.com/jmoiron/sqlx"
)

// Routes adds specific routes for this group.
func Routes(build string, app *web.App, log *logger.Logger, db *sqlx.DB, authClient *authclient.Client) {
	authen := mid.AuthenticateService(log, authClient)
	athAdminOnly := mid.AuthorizeService(log, authClient, auth.RuleAdminOnly)

	api := newAPI(build, log, db)
	app.HandleFuncNoMiddleware("GET /liveness", api.liveness)
	app.HandleFuncNoMiddleware("GET /readiness", api.readiness)
	app.HandleFunc("GET /testerror", api.testError)
	app.HandleFunc("GET /testpanic", api.testPanic)
	app.HandleFunc("GET /testauth", api.liveness, authen, athAdminOnly)
}
