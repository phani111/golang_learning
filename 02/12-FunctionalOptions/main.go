package main

import (
	"12-FunctionalOptions/builder"
	"12-FunctionalOptions/options"
	"log"
)

func main() {

	cfg, err := builder.New().Port(9000).Build()
	if err != nil {
		panic(err)
	}
	log.Println(cfg)

	s, err := builder.NewServer("localhost", cfg)
	if err != nil {
		panic(err)
	}
	log.Println(s)

	s2, err := options.NewServer(
		"localhost",
		options.WithPort(9000),
	)
	if err != nil {
		panic(err)
	}

	log.Println(s2)

}
