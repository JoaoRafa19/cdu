package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Pokemons struct {

}

func main() {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto")
	if err != nil {
		fmt.Print(err.Error())
		panic(err)
	}

	
	
	 responseData, err := io.ReadAll(response.Body)
	 if err != nil {
		fmt.Print(err.Error())
		panic(err)
	}






	fmt.Print(string(responseData))


	var resp Response

	json.Unmarshal(responseData, &resp)

	for _, poke := range resp.Pokemons {
		fmt.Println(poke)
	}
}

type Response struct {
	Nome string `json:"name"`
	Pokemons []Pokemon `json:"pokemon_entries"`

}

type Pokemon struct {
	Number int `json:"entry_number"`
	Especie PokemonSpecies `json:"pokemon_species"`

}

type PokemonSpecies struct {
	Nome string `json:"name"`
}