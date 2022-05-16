package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/felixbecker/hexadiscountexample/api"
	"github.com/felixbecker/hexadiscountexample/application"
	"github.com/felixbecker/hexadiscountexample/web"
)

type fakeDiscounter struct{}

func (f *fakeDiscounter) Rate(amount float32) float32 {

	return 1
}

func main() {

	fmt.Println("Hello http service")

	app := application.New(&fakeDiscounter{})
	ui := web.New(app)
	api := api.New(app)

	server := http.NewServeMux()
	server.Handle("/", ui)
	server.Handle("/api/", http.StripPrefix("/api", api))
	log.Fatal(http.ListenAndServe(":8080", server))
}
