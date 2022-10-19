package withsqlitedb

import (
	"context"
	"demo3/domain_core/model/entity"
	"demo3/shared/driver"
	"demo3/shared/infrastructure/config"
	"demo3/shared/infrastructure/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type gateway struct {
	log     logger.Logger
	appData driver.ApplicationData
	config  *config.Config

	db *gorm.DB
}

// NewGateway ...
func NewGateway(log logger.Logger, appData driver.ApplicationData, cfg *config.Config) *gateway {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	err = db.AutoMigrate(entity.Todo{})
	if err != nil {
		panic(err.Error())
	}

	return &gateway{
		log:     log,
		appData: appData,
		config:  cfg,
		db:      db,
	}
}

func (r *gateway) FindAllTodo(ctx context.Context, page, size int64) ([]*entity.Todo, int64, error) {
	r.log.Info(ctx, "called")

	var count int64
	res := make([]*entity.Todo, 0)
	err := r.db.Model(&entity.Todo{}).Count(&count).Offset(int((page - 1) * size)).Limit(int(size)).Find(&res).Error
	if err != nil {
		return nil, 0, err
	}

	return res, count, nil
}

func (r *gateway) FindOneTodo(ctx context.Context, todoID string) (*entity.Todo, error) {
	r.log.Info(ctx, "called")

	var res entity.Todo
	err := r.db.First(&res, "id = ?", todoID).Error
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (r *gateway) SaveTodo(ctx context.Context, obj *entity.Todo) error {
	r.log.Info(ctx, "called")

	err := r.db.Save(obj).Error
	if err != nil {
		return err
	}

	return nil
}
