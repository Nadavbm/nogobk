package dat

import (
	"database/sql"
	"time"

	"github.com/nadavbm/nogobk/pkg/logger"
)

type Session struct {
	UserId  int64     `db:"userid"`
	Created time.Time `db:"created"`
	Token   string    `db:"token"`
	Csrf    string    `db:"csrf"`
	Expires time.Time `db:"expires"`
}

type SessionMapper struct {
	Txn *sql.Tx
}

func (s *Session) CreateSession(l *logger.Logger) error {
	conn := GetDBConnString()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		l.Panic("could not open connection to database")
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO sessions(userid, created, token, csrf, expires) VALUES ($1,$2,$3,$4,$5)", s.UserId, time.Now(), s.Token, s.Csrf, s.Expires)

	return err
}

func (s *Session) GetSessionByToken(l *logger.Logger) error {
	conn := GetDBConnString()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		l.Panic("could not open connection to database")
	}
	defer db.Close()

	_, err = db.Exec("SELECT created, token, csrf, expires FROM sessions WHERE userid = $1 and token = $2", s.UserId, s.Token)

	return err
}
