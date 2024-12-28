package db

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Store interface {
	Querier
	AddToCartTx(ctx context.Context, arg CreateCartitemParams) (CartTxResult, error)
	UpdateCartTx(ctx context.Context, cartItemID uuid.UUID, arg UpdateCartitemParams) (CartTxResult, error)
	RemoveCartTx(ctx context.Context, cartItemID uuid.UUID) (RemoveCartTxResult, error)
}

type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

func (store *SQLStore) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}
