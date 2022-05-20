package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/felixbecker/hexadiscountexample/api"
	"github.com/felixbecker/hexadiscountexample/application"
	internal "github.com/felixbecker/hexadiscountexample/internal/factory"
	"github.com/felixbecker/hexadiscountexample/web"
)

func main() {

	fmt.Println("Hello http service")

	config := internal.Config{
		StoreType: "inmemory",
		Redis: internal.Redis{
			Addr: "localhost:6379",
		},
	}
	factory := internal.NewFactory(&config)
	app := application.New(factory.Discounter())
	ui := web.New(app)
	api := api.New(app)

	server := http.NewServeMux()
	server.Handle("/", ui)
	server.Handle("/api/", http.StripPrefix("/api", api))
	log.Fatal(http.ListenAndServe(":8080", server))
}
