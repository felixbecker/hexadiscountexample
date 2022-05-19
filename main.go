package main

import (
	"fmt"

	"github.com/felixbecker/hexadiscountexample/storeprovider"
)

func main() {

	fmt.Println("redis")

	p := storeprovider.NewRedisProvider("localhost:6379")
	rate := p.Get(1.0)

	fmt.Println(rate)
	p.Set(1.24, 2.23)
	rate = p.Get(1.0)
	fmt.Println(rate)
}
