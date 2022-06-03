package main

import (
	"log"
	"os"

	"github.com/felixbecker/hexadiscountexample/cli"
	internal "github.com/felixbecker/hexadiscountexample/internal/factory"
)

func main() {

	factory, err := internal.NewFactory()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	commandLine := cli.New(factory.Application())
	commandLine.Execute()
}
