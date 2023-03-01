package jason

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PokemonType struct {
	Name string `json:"name"`
}

type Type struct {
	PokemonType `json:"type"`
}

type PokemonForm struct {
	Height int    `json:"height"`
	Weight int    `json:"weight"`
	Types  []Type `json:"types"`
}

type PokemonStats struct {
	Height int
	Weight int
	Types  []string
}

const url = "https://pokeapi.co/api/v2/pokemon"

func GetPokemonStats(name string) (*PokemonStats, error) {
	pf := &PokemonForm{}
	res, err := http.Get(fmt.Sprintf("%s/%s", url, name))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(pf)
	if err != nil {
		return nil, err
	}
	var types []string
	for _, t := range pf.Types {
		types = append(types, t.PokemonType.Name)
	}
	ps := &PokemonStats{
		Height: pf.Height,
		Weight: pf.Weight,
		Types:  types,
	}
	return ps, nil
}
