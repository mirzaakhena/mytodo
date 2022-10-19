package repository

import (
	"context"
	"demo3/domain_core/model/entity"
)

type SaveTodoRepo interface {
	SaveTodo(ctx context.Context, obj *entity.Todo) error
}

type FindOneTodoRepo interface {
	FindOneTodo(ctx context.Context, todoID string) (*entity.Todo, error)
}

type FindAllTodoRepo interface {
	FindAllTodo(ctx context.Context, page, size int64) ([]*entity.Todo, int64, error)
}
