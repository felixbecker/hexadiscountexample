package main

import (
	"os"

	"github.com/felixbecker/hexadiscountexample/cli"
	internal "github.com/felixbecker/hexadiscountexample/internal/factory"
)

func main() {

	dbUser := os.Getenv("DB_USER")
	dbPW := os.Getenv("DB_PW")
	config := internal.Config{
		StoreType: "postgres",
		Redis: internal.Redis{
			Addr: "localhost:6379",
		},
		Postgres: internal.Postgres{
			User:      dbUser,
			Password:  dbPW,
			Host:      "localhost:5432",
			DB:        "discounter",
			Tablename: "rates",
		},
	}

	factory := internal.NewFactory(&config)
	commandLine := cli.New(factory.Application())
	commandLine.Execute()
}
