package dat

import (
	"database/sql"
	"time"
)

type Session struct {
	UserId  int64     `db:"userid"`
	Created time.Time `db:"created"`
	Token   string    `db:"token"`
	Csrf    string    `db:"csrf"`
	Expires time.Time `db:"expires"`
}

type SessionMapper struct {
	txn *sql.Tx
}

func CreateSession() {

}

func GetSessionByToken() {

}
