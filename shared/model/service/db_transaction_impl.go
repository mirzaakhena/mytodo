package service

import (
	"context"
	"demo3/shared/model/repository"
)

// WithoutTransaction is helper function that simplify the readonly db
func WithoutTransaction[T any](ctx context.Context, trx repository.WithoutTransactionDB, trxFunc func(dbCtx context.Context) (T, error)) (T, error) {
	dbCtx, err := trx.GetDatabase(ctx)
	if err != nil {
		return nil, err
	}

	defer func(trx repository.WithoutTransactionDB, ctx context.Context) {
		err := trx.Close(ctx)
		if err != nil {
			return
		}
	}(trx, dbCtx)

	return trxFunc(dbCtx)
}

// WithTransaction is helper function that simplify the transaction execution handling
func WithTransaction[T any](ctx context.Context, trx repository.WithTransactionDB, trxFunc func(dbCtx context.Context) (T, error)) (T, error) {
	dbCtx, err := trx.BeginTransaction(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if p := recover(); p != nil {
			err = trx.RollbackTransaction(dbCtx)
			panic(p)

		} else if err != nil {
			err = trx.RollbackTransaction(dbCtx)

		} else {
			err = trx.CommitTransaction(dbCtx)

		}
	}()

	return trxFunc(dbCtx)
}
