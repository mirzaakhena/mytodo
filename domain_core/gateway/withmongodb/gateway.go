package withmongodb

import (
	"context"
	"demo3/domain_core/model/entity"
	"demo3/shared/driver"
	"demo3/shared/infrastructure/config"
	"demo3/shared/infrastructure/database"
	"demo3/shared/infrastructure/logger"
)

type gateway struct {
	log     logger.Logger
	appData driver.ApplicationData
	config  *config.Config
	repo    database.Repository[entity.Todo]
}

// NewGateway ...
func NewGateway(log logger.Logger, appData driver.ApplicationData, cfg *config.Config) *gateway {

	db := database.NewDatabase("todo_app_db")

	return &gateway{
		log:     log,
		appData: appData,
		config:  cfg,
		repo:    database.NewMongoGateway[entity.Todo](db),
	}
}

func (r *gateway) FindAllTodo(ctx context.Context, page, size int64) ([]*entity.Todo, int64, error) {
	r.log.Info(ctx, "called")

	param := database.NewDefaultParam().
		SetPage(page).
		SetSize(size)

	res := make([]*entity.Todo, 0)
	count, err := r.repo.GetAll(param, &res)
	if err != nil {
		return nil, 0, err
	}

	return res, count, nil
}

func (r *gateway) FindOneTodo(ctx context.Context, todoID string) (*entity.Todo, error) {
	r.log.Info(ctx, "called")

	filter := map[string]any{
		"_id": todoID,
	}

	var res entity.Todo
	err := r.repo.GetOne(filter, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *gateway) SaveTodo(ctx context.Context, obj *entity.Todo) error {
	r.log.Info(ctx, "called")

	err := r.repo.InsertOrUpdate(obj)
	if err != nil {
		return err
	}

	return nil
}
