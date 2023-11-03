package jason_test

import (
	"github.com/cpustejovsky/ready-steady-go/tdd/jason"
	"reflect"
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
	if err != nil {
		t.Fatalf("got:\t%v\nwant:\t%v\n", err, nil)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got:\t%v\nwant:\t%v\n", got, want)
	}
}
