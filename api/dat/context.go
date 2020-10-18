package dat

import (
	"database/sql"
)

type Context struct {
	Txn      *sql.Tx
	Users    *UserMapper
	Sessions *SessionMapper
}

func NewDatabaseContext() (*Context, error) {
	txn, err := db.Begin()
	if err != nil {
		return nil, err
	}

	c := Context{
		Txn:      txn,
		Users:    &UserMapper{txn: txn},
		Sessions: &SessionMapper{txn: txn},
	}
	return &c, nil
}
