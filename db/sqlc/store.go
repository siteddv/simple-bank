package db

import (
	"context"
	"database/sql"
	"fmt"
)

type Store struct {
	queries *Queries
	db      *sql.DB
}

func NewStore(db *sql.DB) *Store {
	res := &Store{
		db:      db,
		queries: New(db),
	}
	return res
}

// execTr executes a function within database transaction
func (store *Store) execTr(ctx context.Context,
	fn func(*Queries) error) error {

	tx, err := store.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	queries := New(tx)
	if err = fn(queries); err != nil {
		if rlError := tx.Rollback(); rlError != nil {
			return fmt.Errorf("transaction err: %w, rollback error: %v", err, rlError)
		}
		return err
	}

	err = tx.Commit()
	return err
}

// TransferTxParams contains the input parameters of the transfer transaction
type TransferTxParams struct {
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

// TransferTxResult is the result of transfer transaction
type TransferTxResult struct {
	Transfer    Transfer `json:"transfer"`
	FromAccount Account  `json:"from_account"`
	ToAccount   Account  `json:"to_account"`
	FromEntry   Entry    `json:"from_entry"`
	ToEntry     Entry    `json:"to_entry"`
}

// TransferTx performs a money transfer from one account to other
// It creates a transfer record, add account entries, and update accounts' balance within a single database transaction
func (store *Store) TransferTx(ctx context.Context, arg TransferTxParams) (TransferTxResult, error) {
	var result TransferTxResult

	err := store.execTr(ctx, func(q *Queries) error {
		var execError error

		createTransferArgs := CreateTransferParams{
			FromAccountID: arg.FromAccountID,
			ToAccountID:   arg.ToAccountID,
			Amount:        arg.Amount,
		}
		result.Transfer, execError = q.CreateTransfer(ctx, createTransferArgs)
		if execError != nil {
			return execError
		}

		createEntryArgs := CreateEntryParams{
			AccountID: arg.FromAccountID,
			Amount:    -arg.Amount,
		}
		result.FromEntry, execError = q.CreateEntry(ctx, createEntryArgs)
		if execError != nil {
			return execError
		}

		createEntryArgs = CreateEntryParams{
			AccountID: arg.ToAccountID,
			Amount:    arg.Amount,
		}
		result.ToEntry, execError = q.CreateEntry(ctx, createEntryArgs)
		if execError != nil {
			return execError
		}

		result.FromAccount, execError = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.FromAccountID,
			Amount: -arg.Amount,
		})
		if execError != nil {
			return execError
		}

		result.ToAccount, execError = q.AddAccountBalance(ctx, AddAccountBalanceParams{
			ID:     arg.ToAccountID,
			Amount: arg.Amount,
		})

		return execError
	})

	return result, err
}
