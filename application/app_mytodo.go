package application

import (
	"demo3/domain_core/controller/restapi"
	"demo3/domain_core/gateway/withsqlitedb"
	"demo3/domain_core/usecase/getalltodo"
	"demo3/domain_core/usecase/runtodocheck"
	"demo3/domain_core/usecase/runtodocreate"
	"demo3/shared/driver"
	"demo3/shared/infrastructure/config"
	"demo3/shared/infrastructure/logger"
	"demo3/shared/infrastructure/server"
	"demo3/shared/util"
)

type mytodo struct {
	httpHandler *server.GinHTTPHandler
	controller  driver.Controller
}

func (c mytodo) RunApplication() {
	c.controller.RegisterRouter()
	c.httpHandler.RunApplication()
}

func NewMytodo() func() driver.RegistryContract {

	const appName = "mytodo"

	return func() driver.RegistryContract {

		cfg := config.ReadConfig()

		appID := util.GenerateID(4)

		appData := driver.NewApplicationData(appName, appID)

		log := logger.NewSimpleJSONLogger(appData)

		httpHandler := server.NewGinHTTPHandler(log, cfg.Servers[appName].Address, appData)

		datasource := withsqlitedb.NewGateway(log, appData, cfg)

		return &mytodo{
			httpHandler: &httpHandler,
			controller: &restapi.Controller{
				Log:                 log,
				Config:              cfg,
				Router:              httpHandler.Router,
				GetAllTodoInport:    getalltodo.NewUsecase(datasource),
				RunTodoCheckInport:  runtodocheck.NewUsecase(datasource),
				RunTodoCreateInport: runtodocreate.NewUsecase(datasource),
			},
		}

	}
}
