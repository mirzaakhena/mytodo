package runtodocheck

import (
	"context"

	"demo3/shared/usecase"
)

type Inport usecase.Inport[context.Context, InportRequest, InportResponse]

// InportRequest is request payload to run the usecase
type InportRequest struct {
	TodoID string
}

// InportResponse is response payload after running the usecase
type InportResponse struct {
}
