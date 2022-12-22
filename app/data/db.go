package data

import (
	"fmt"

	"github.com/caseyrwebb/auth-microservice/app/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

type AuthDB interface {
}

type GoDB interface {
	AuthDB
	Open() error
	Close() error
}

type DB struct {
	db     *sqlx.DB
	logger *zap.Logger
}

func (d *DB) SetDBLogger(logger *zap.Logger) {
	d.logger = logger
}

func (d *DB) Open(config *utils.Configurations) error {
	var conn string

	if config.DBConn != "" {
		conn = config.DBConn
	} else {
		host := config.DBHost
		port := config.DBPort
		user := config.DBUser
		dbName := config.DBName
		password := config.DBPass
		conn = fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", host, port, user, dbName, password)
	}
	d.logger.Debug(fmt.Sprintf("%s %s", "connection string", conn))

	pg, err := sqlx.Open("postgres", conn)
	if err != nil {
		return err
	}
	d.logger.Debug("Connected to Database!")

	pg.MustExec(createUserTableSchema)
	pg.MustExec(createVerificationSchema)

	d.db = pg

	return nil
}

func (d *DB) Close() error {
	return d.db.Close()
}
