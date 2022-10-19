package getalltodo

import (
	"context"
	"demo3/shared/util"
)

//go:generate mockery --name Outport -output mocks/

type getAllTodoInteractor struct {
	outport Outport
}

// NewUsecase is constructor for create default implementation of usecase
func NewUsecase(outputPort Outport) Inport {
	return &getAllTodoInteractor{
		outport: outputPort,
	}
}

// Execute the usecase
func (r *getAllTodoInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

	res := &InportResponse{}

	todoObjs, count, err := r.outport.FindAllTodo(ctx, req.Page, req.Size)
	if err != nil {
		return nil, err
	}

	res.Count = count
	res.Items = util.ToSliceAny(todoObjs)

	return res, nil
}
