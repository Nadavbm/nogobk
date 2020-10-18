package dat

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/nadavbm/nogobk/pkg/logger"
	"go.uber.org/zap"
)

type User struct {
	Id         int64     `db:"id"`
	Created    time.Time `db:"created"`
	Name       string    `db:"name"`
	Email      string    `db:"email"`
	Password   string    `db:"paswword"`
	AuthMethod string    `db:"auth_method"`
}

type UserMapper struct {
	txn *sql.Tx
}

// CreateUser will create user in the database from a request handler in /signup
func (m *UserMapper) CreateUser(l *logger.Logger, n, e, p string) (*User, error) {
	u := &User{
		Name:     n,
		Email:    e,
		Password: p,
	}

	query := fmt.Sprintf("INSERT INTO users(name, email, password) VALUES ('%s', '%s', '%s');", u.Name, u.Email, u.Password)
	_, err := m.txn.Query(query)
	if err != nil {
		l.Info("ERROR:", zap.Error(err))
		l.Info("could not execute in database:", zap.String("query:", query))
	}
	l.Info("execute in database:", zap.String("query:", query))
	return u, nil
}

// GetUserByEmail will be used by /login to get to user login credentials
func (m *UserMapper) GetUserByEmail(l *logger.Logger, email, password string) (*User, error) {
	user := &User{}
	rows, err := m.txn.Query("SELECT name,email FROM users WHERE email = '%s'", email)

	//query := fmt.Sprintf("SELECT name,email FROM users WHERE email = '%s'", email)
	//rowss, err := db.Query(query)

	if err != nil {
		l.Info("could not get login credentials from database", zap.Error(err))

	}

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Password)
		if err != nil {
			l.Info("could not scan users table")
		}
	}

	fmt.Println("from database:", user.Email, user.Password)
	return user, nil
}

// GetUserById will be used by /profile{user} to get to user profile
func (m *UserMapper) GetUserById(l *logger.Logger, id int64) (*User, error) {
	user := &User{}
	rows, err := m.txn.Query("SELECT name,email FROM users WHERE id = %d", id)

	//query := fmt.Sprintf("SELECT name,email FROM users WHERE email = '%s'", id)
	//rowss, err := db.Query(query)
	if err != nil {
		l.Info("could not get login credentials from database", zap.Error(err))

	}

	//expiresAt := time.Now().Add(time.Minute * 100000).Unix()

	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Password)
		if err != nil {
			l.Info("could not scan users table")
		}
	}

	fmt.Println("from database:", user.Email, user.Password)
	return user, nil
}
