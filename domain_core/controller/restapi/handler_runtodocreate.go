package restapi

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"demo3/domain_core/usecase/runtodocreate"
	"demo3/shared/infrastructure/logger"
	"demo3/shared/model/payload"
	"demo3/shared/util"
)

// runTodoCreateHandler ...
func (r *Controller) runTodoCreateHandler() gin.HandlerFunc {

	type request struct {
		runtodocreate.InportRequest
	}

	type response struct {
	}

	return func(c *gin.Context) {

		traceID := util.GenerateID(16)

		ctx := logger.SetTraceID(context.Background(), traceID)

		var jsonReq request
		if err := c.BindJSON(&jsonReq); err != nil {
			r.Log.Error(ctx, err.Error())
			c.JSON(http.StatusBadRequest, payload.NewErrorResponse(err, traceID))
			return
		}

		var req runtodocreate.InportRequest
		req.TodoRequest = jsonReq.TodoRequest
		req.Now = time.Now()
		req.RandomString = util.GenerateID(12)

		r.Log.Info(ctx, util.MustJSON(req))

		res, err := r.RunTodoCreateInport.Execute(ctx, req)
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
