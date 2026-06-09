package mux

import (
	"os"

	"github.com/DimaMaimesko/ultimate-go-service/api/services/sales/route/sys/checkapi"
	"github.com/DimaMaimesko/ultimate-go-service/foundation/web"
)

// WebAPI constructs a http.Handler with all application routes bound.
func WebAPI(shutdown chan os.Signal) *web.App {
	mux := web.NewApp(shutdown)

	checkapi.Routes(mux)

	return mux
}
