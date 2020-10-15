package config

import (
	"context"
	"fmt"
	"os"

	"github.com/go-pg/pg/v10"
	log "github.com/sirupsen/logrus"
)

func ReadPostgres() (*pg.DB, error) {
	return CreateDBConnectionPostgres(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("READ_DB_USER_POSTGRES"),
		os.Getenv("READ_DB_PASSWORD_POSTGRES"),
		os.Getenv("READ_DB_HOST_POSTGRES"),
		os.Getenv("READ_DB_PORT_POSTGRES"),
		os.Getenv("READ_DB_NAME_POSTGRES")))
}

func WritePostgres() (*pg.DB, error) {
	return CreateDBConnectionPostgres(fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("READ_DB_USER_POSTGRES"),
		os.Getenv("READ_DB_PASSWORD_POSTGRES"),
		os.Getenv("READ_DB_HOST_POSTGRES"),
		os.Getenv("READ_DB_PORT_POSTGRES"),
		os.Getenv("READ_DB_NAME_POSTGRES")))
}

// CreateDBConnection function for creating database connection
func CreateDBConnectionPostgres(url string) (*pg.DB, error) {

	opt, err := pg.ParseURL(url)
	if err != nil {
		log.Info("Invalid database URL!")
		log.Fatal(err)
	}

	opt.MaxConnAge = 100
	opt.IdleTimeout = 5
	opt.PoolSize = 50

	client := pg.Connect(opt)

	err = client.Ping(context.TODO())
	if err != nil {
		log.Info("not connect database")
		log.Fatal(err)
		defer client.Close()
	}

	log.Info("Connected to Postgres!")
	return client, err

}
