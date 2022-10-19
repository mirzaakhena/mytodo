package restapi

import (
	"github.com/gin-gonic/gin"

	"demo3/domain_core/usecase/getalltodo"
	"demo3/domain_core/usecase/runtodocheck"
	"demo3/domain_core/usecase/runtodocreate"
	"demo3/shared/infrastructure/config"
	"demo3/shared/infrastructure/logger"
)

type Controller struct {
	Router              gin.IRouter
	Config              *config.Config
	Log                 logger.Logger
	GetAllTodoInport    getalltodo.Inport
	RunTodoCheckInport  runtodocheck.Inport
	RunTodoCreateInport runtodocreate.Inport
}

// RegisterRouter registering all the router
func (r *Controller) RegisterRouter() {
	r.Router.GET("/todo", r.authorized(), r.getAllTodoHandler())
	r.Router.PUT("/todo/:todo_id", r.authorized(), r.runTodoCheckHandler())
	r.Router.POST("/todo", r.authorized(), r.runTodoCreateHandler())
}
