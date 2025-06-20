package main

import "github.com/r2adio/pokedex/internal/pokeapi"

// struct that holds all the stateful info for command callback functions
type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

func main() {
	cfg := config{pokeapiClient: pokeapi.NewClient()}

	startREPL(&cfg)
}
