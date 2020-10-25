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
	Password   string    `db:"password"`
	AuthMethod string    `db:"auth_method"`
}

type UserMapper struct {
	Txn *sql.Tx
}

// methods using user struct directly

func (u *User) CreateUsers(l *logger.Logger) error {
	conn := GetDBConnString()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		l.Panic("could not open connection to database")
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO users(name, created, email, password, auth_method) VALUES ($1,$2,$3,$4,$5)", u.Name, time.Now(), u.Email, u.Password, "basic_auth")
	return err
}

func GetUserByEmail(l *logger.Logger, email string) (*User, error) {
	conn := GetDBConnString()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		l.Panic("could not open connection to database")
	}
	defer db.Close()

	query := fmt.Sprintf("SELECT id,email, password FROM users WHERE email = '%s'", email)
	rows, err := db.Query(query)
	if err != nil {
		l.Info("could not get login credentials from database", zap.Error(err))
		return nil, err
	}

	user := new(User)
	for rows.Next() {
		err := rows.Scan(&user.Id, &user.Email, &user.Password)
		if err != nil {
			l.Info("could not scan users table")
		}
	}

	return user, nil
}

// Methods using db mapper

// CreateUser will create user in the database from a request handler in /signup
func (m *UserMapper) CreateUser(l *logger.Logger, n, e, p string) (*User, error) {
	fmt.Println("beginning of user mapper:", m, m.Txn)
	u := &User{
		Name:     n,
		Email:    e,
		Password: p,
	}

	query := fmt.Sprintf("INSERT INTO users(name, created, email, password, auth_method) VALUES ($1,$2,$3,$4,$5);", u.Name, time.Now(), u.Email, u.Password, "basic_auth")
	_, err := m.Txn.Exec("INSERT INTO users(name, created, email, password, auth_method) VALUES ($1,$2,$3,$4,$5)")
	if err != nil {
		l.Info("ERROR:", zap.Error(err))
		l.Info("could not execute in database:", zap.String("query:", query))
	}
	l.Info("execute in database:", zap.String("query:", query))
	return u, nil
}

// GetUserByEmail will be used by /login to get to user login credentials
func (m *UserMapper) GetUserByEmailm(l *logger.Logger, email, password string) (*User, error) {
	user := &User{}
	rows, err := m.Txn.Query("SELECT name,email FROM users WHERE email = '%s'", email)

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
	rows, err := m.Txn.Query("SELECT name,email FROM users WHERE id = %d", id)

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
