package gensql

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (s *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	//todo: make this code little bit clear;
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := New(tx)
	err = fn(q)
	if err != nil {
		rbError := tx.Rollback()
		if rbError != nil {
			return fmt.Errorf("tx err %v, rollback error: %v", err, rbError)
		}

		return err
	}

	return tx.Commit()
}

type TransferTransactionArgs struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTransactionResult struct {
	Transfer    Transfer `json:"tranfer"`
	FromAccount Account  `json:"to_account"`
	ToAccocunt  Account  `json:"from_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

func (s *Store) TransferTx(ctx context.Context, arg TransferTransactionArgs) (TransferTransactionResult, error) {
	var result TransferTransactionResult
	err := s.execTx(ctx, func(q *Queries) error {
		var err error
		var fromAccount Account
		var toAccount Account
		result.Transfer, err = q.CreateTransfer(ctx, CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		})
		if err != nil {
			return err
		}
		//note: from enttiy record
		result.FromEntry, err = q.CreateEntity(ctx, CreateEntityParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		})

		if err != nil {
			return err
		}

		//note:to entity record
		result.ToEntry, err = q.CreateEntity(ctx, CreateEntityParams{
			AccountID: arg.FromAccountID,
			Amount:    arg.Amount,
		})
		if err != nil {
			return err
		}

		//note: update acconts
		fromAccount, err = q.GetAccount(ctx, arg.FromAccountID)
		if err != nil {
			return err
		}
		toAccount, err = q.GetAccount(ctx, arg.ToAccountID)
		if err != nil {
			return err
		}

		result.FromAccount, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.FromAccountID,
			Balance: fromAccount.Balance - arg.Amount,
		})
		if err != nil {
			return err
		}

		result.ToAccocunt, err = q.UpdateAccount(ctx, UpdateAccountParams{
			ID:      arg.ToAccountID,
			Balance: toAccount.Balance + arg.Amount,
		})

		if err != nil {
			return err
		}
		return err
	})

	return result, err

}
