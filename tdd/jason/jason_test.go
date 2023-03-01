package jason_test

import (
	"github.com/cpustejovsky/ready-steady-go/tdd/jason"
	"github.com/stretchr/testify/assert"
	"testing"
)

//Goal: Get me a pikachu's basic stats: height weight type

func TestGetPokemonStats(t *testing.T) {
	want := &jason.PokemonStats{
		Height: 4,
		Weight: 60,
		Types:  []string{"electric"},
	}
	got, err := jason.GetPokemonStats("pikachu")
	assert.Nil(t, err)
	assert.Equal(t, want, got)
}
