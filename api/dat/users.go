package dat

import (
	"database/sql/driver"
	"time"
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
	txn *driver.Tx
}

func CreateUser() {

}

func GetUserById() {

}
