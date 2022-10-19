package runtodocreate

import (
	"context"
	"demo3/domain_core/model/entity"
	"demo3/shared/usecase"
)

type Inport usecase.Inport[context.Context, InportRequest, InportResponse]

// InportRequest is request payload to run the usecase
type InportRequest struct {
	entity.TodoRequest
}

// InportResponse is response payload after running the usecase
type InportResponse struct {
}
