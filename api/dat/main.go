package dat

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/nadavbm/nogobk/pkg/env"
	"github.com/nadavbm/nogobk/pkg/logger"
	"go.uber.org/zap"
)

var db *sql.DB

func InitDB() {
	logger := logger.SetLogger()

	conn := GetDBConnString()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		logger.Panic("could not open connection to database", zap.Error(err))
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Panic("could not ping database")
	}
	logger.Info("connected to database: " + env.DatabaseDB + "on host:" + env.DatabaseHost)

	_, err = db.Exec(migration)
	if err != nil {
		logger.Panic("could not run db migrations", zap.Error(err))
	}
	logger.Info("db migration completed")

	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	_, err = tx.ExecContext(ctx, "INSERT INTO pets (name, species) VALUES ('Fido', 'dog'), ('Albert', 'cat')")
	if err != nil {
		// Incase we find any error in the query execution, rollback the transaction
		tx.Rollback()
		return
	}
}

func GetDBConnString() string {
	conn := fmt.Sprintf("postgres://%s:%s@%s:%v/%s?sslmode=disable", env.DatabaseUser, env.DatabasePass, env.DatabaseHost, env.DatabasePort, env.DatabaseDB)
	return conn
}

func getDBConnStr() string {
	conn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable", env.DatabaseHost, env.DatabasePort, env.DatabaseUser, env.DatabasePass, env.DatabaseDB)
	return conn
}
