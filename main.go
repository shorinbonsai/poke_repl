package main

import (
	"os"
	"poke_repl/internal/pokeapi"
	"time"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour / 2),
	}
	Start(&cfg, os.Stdin, os.Stdout)
}
