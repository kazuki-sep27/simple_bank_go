// Code generated by sqlc. DO NOT EDIT.
// source: account.sql

package db

import (
	"context"
	"database/sql"
)

const createAccounts = `-- name: CreateAccounts :execresult
INSERT INTO accounts (
  owner,balance,currency
) VALUES (
  ?, ?, ?
)
`

type CreateAccountsParams struct {
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

func (q *Queries) CreateAccounts(ctx context.Context, arg CreateAccountsParams) (sql.Result, error) {
	return q.exec(ctx, q.createAccountsStmt, createAccounts, arg.Owner, arg.Balance, arg.Currency)
}

const deleteAccount = `-- name: DeleteAccount :exec
DELETE FROM accounts
WHERE id = ?
`

func (q *Queries) DeleteAccount(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteAccountStmt, deleteAccount, id)
	return err
}

const getAccountByID = `-- name: GetAccountByID :one
SELECT id,owner,balance,currency FROM accounts
WHERE id = ?
`

type GetAccountByIDRow struct {
	ID       int64  `json:"id"`
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

func (q *Queries) GetAccountByID(ctx context.Context, id int64) (GetAccountByIDRow, error) {
	row := q.queryRow(ctx, q.getAccountByIDStmt, getAccountByID, id)
	var i GetAccountByIDRow
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
	)
	return i, err
}

const getLastAccount = `-- name: GetLastAccount :one
SELECT id,owner,balance,currency FROM accounts
ORDER BY id DESC 
LIMIT 1
`

type GetLastAccountRow struct {
	ID       int64  `json:"id"`
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

func (q *Queries) GetLastAccount(ctx context.Context) (GetLastAccountRow, error) {
	row := q.queryRow(ctx, q.getLastAccountStmt, getLastAccount)
	var i GetLastAccountRow
	err := row.Scan(
		&i.ID,
		&i.Owner,
		&i.Balance,
		&i.Currency,
	)
	return i, err
}

const listAccounts = `-- name: ListAccounts :many
SELECT id,owner,balance,currency FROM accounts
ORDER BY id
`

type ListAccountsRow struct {
	ID       int64  `json:"id"`
	Owner    string `json:"owner"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}

func (q *Queries) ListAccounts(ctx context.Context) ([]ListAccountsRow, error) {
	rows, err := q.query(ctx, q.listAccountsStmt, listAccounts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListAccountsRow
	for rows.Next() {
		var i ListAccountsRow
		if err := rows.Scan(
			&i.ID,
			&i.Owner,
			&i.Balance,
			&i.Currency,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
