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
	caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour / 2),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	Start(&cfg, os.Stdin, os.Stdout)
}
