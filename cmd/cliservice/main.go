package main

import (
	"github.com/felixbecker/hexadiscountexample/application"
	"github.com/felixbecker/hexadiscountexample/cli"
)

type fakeDiscounter struct {
}

func (f *fakeDiscounter) Rate(amount float32) float32 {
	return 1
}

func main() {

	app := application.New(&fakeDiscounter{})
	commandLine := cli.New(app)
	commandLine.Execute()
}
