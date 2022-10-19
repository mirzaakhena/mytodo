package restapi

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"

	"demo3/domain_core/usecase/runtodocheck"
	"demo3/shared/infrastructure/logger"
	"demo3/shared/model/payload"
	"demo3/shared/util"
)

// runTodoCheckHandler ...
func (r *Controller) runTodoCheckHandler() gin.HandlerFunc {

	type request struct {
	}

	type response struct {
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		//var jsonReq request
		//if err := c.BindJSON(&jsonReq); err != nil {
		//	r.Log.Error(ctx, err.Error())
		//	c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
		//	return
		//}

		var req runtodocheck.InportRequest
		req.TodoID = c.Param("todo_id")

		r.Log.Info(ctx, util.MustJSON(req))

		res, err := r.RunTodoCheckInport.Execute(ctx, req)
		if err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var jsonRes response
		_ = res

		r.Log.Info(ctx, util.MustJSON(jsonRes))
		c.JSON(http.StatusOK, payload.NewSuccessResponse(jsonRes, traceID))

	}
}
