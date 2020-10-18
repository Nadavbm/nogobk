package dat

import (
	"context"
	"database/sql"
)

type DbMap struct {
	Db *sql.DB
}

type Transaction struct {
	dbmap *DbMap
	txn   *sql.Tx
}

func (m *DbMap) BeginTxn() (*Transaction, error) {
	ctx := context.Background()
	txn, err := m.Db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}

	t := &Transaction{
		txn: txn,
	}
	return t, nil
}
