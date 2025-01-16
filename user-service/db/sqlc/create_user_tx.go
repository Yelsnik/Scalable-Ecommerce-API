package db

import (
	"context"
	"fmt"
)

type CreateUserTxParams struct {
	CreateUserParams
	AfterCreate func(user User) error
}

type CreateUserTxResult struct {
	User User
}

func (store *SQLStore) CreateUserTx(ctx context.Context, arg CreateUserTxParams) (CreateUserTxResult, error) {
	var result CreateUserTxResult

	err := store.execTx(ctx, func(q *Queries) error {
		var err error

		result.User, err = q.CreateUser(ctx, arg.CreateUserParams)
		if err != nil {
			return err
		}

		fmt.Println(result.User)
		err = arg.AfterCreate(result.User)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}
