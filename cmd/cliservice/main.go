package main

import (
	"fmt"

	"github.com/felixbecker/hexadiscountexample/cli"
	internal "github.com/felixbecker/hexadiscountexample/internal/factory"
)

func main() {

	config := internal.Config{
		StoreType: "postgres",
		Redis: internal.Redis{
			Addr: "localhost:6379",
		},
		Postgres: internal.Postgres{
			User:      "root",
			Password:  "root",
			Host:      "localhost:5432",
			DB:        "discounter",
			Tablename: "rates",
		},
	}
	factory := internal.NewFactory(&config)

	fmt.Printf("\n%v\n", factory.Application())
	commandLine := cli.New(factory.Application())
	commandLine.Execute()
}
