package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/felixbecker/hexadiscountexample/api"
	"github.com/felixbecker/hexadiscountexample/application"
	internal "github.com/felixbecker/hexadiscountexample/internal/factory"
	"github.com/felixbecker/hexadiscountexample/web"
)

func main() {

	fmt.Println("Hello http service")

	factory, err := internal.NewFactory()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	app := application.New(factory.Discounter())
	log.Println("Load Hexgon Application")
	ui := web.New(app)
	log.Println("Load Web UI")
	api := api.New(app)
	log.Println("Load API Endpoint")

	server := http.NewServeMux()
	server.Handle("/", ui)
	server.Handle("/api/", http.StripPrefix("/api", api))
	log.Println("Listen and Serve on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", server))
}
