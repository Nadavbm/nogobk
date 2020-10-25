package dat

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/nadavbm/nogobk/pkg/logger"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id         int64     `db:"id"`
	Created    time.Time `db:"created"`
	Name       string    `db:"name"`
	Email      string    `db:"email"`
	Password   string    `db:"password"`
	AuthMethod string    `db:"auth_method"`
}

// methods using user struct directly
func (u *User) Authenticate(l *logger.Logger, email, password string) error {
	fmt.Println("user in db authenticate:", u, "email:", email, "password:", password)
	if email == "" && password == "" {
		l.Info("no email or password provided. enter your credentials please.")
	}

	/*
		emailScan := 0
		rows := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = $1;", email)
		err := rows.Scan(&emailScan)
		if emailScan == 0 || err == sql.ErrNoRows {
			l.Info("email was not found in the database. redirecting to login page.", zap.Error(err))
			return err
		}
	*/

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil || err == bcrypt.ErrMismatchedHashAndPassword {
		l.Info("incorrect password", zap.String("for email:", u.Email))
		return err
	}
	return nil
}

func (u *User) CreateUsers(l *logger.Logger) error {
	conn := GetDBConnString()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		l.Panic("could not open connection to database")
	}
	defer db.Close()

	pass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		l.Info("could not generate passowrd in bcrypt", zap.Error(err))
	}
	u.Password = string(pass)

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
