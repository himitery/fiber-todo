package database

import (
	"context"
	"fmt"
	"log"

	"github.com/himitery/fiber-todo/config"
	"github.com/himitery/fiber-todo/internal/adapter/persistence/sql"
	"github.com/jackc/pgx/v5"
)

type Database struct {
	Queries *sql.Queries
	Context context.Context
}

func NewDatabase(conf *config.Config) *Database {
	ctx := context.Background()

	conn, err := pgx.Connect(
		ctx,
		fmt.Sprintf(
			"postgresql://%s:%s@%s:%s/%s",
			conf.Database.User,
			conf.Database.Password,
			conf.Database.Host,
			conf.Database.Port,
			conf.Database.Database,
		),
	)
	if err != nil {
		log.Fatal(err)
	}

	return &Database{
		Queries: sql.New(conn),
		Context: ctx,
	}
}
