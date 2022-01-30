package server

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) (int, error)
	RecordWin(name string)
}

func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{map[string]int{}}
}

type InMemoryPlayerStore struct {
	Store map[string]int
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) (int, error) {
	score, found := i.Store[name]
	if !found{
		return 0, errors.New("could not find player")
	}
	return score, nil

}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.Store[name]++
}

type PlayerServer struct {
	Store PlayerStore
}

//server.go
func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}

}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score, err := p.Store.GetPlayerScore(player)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

// func RecordWin(name string){

// }

func GetPlayerScore(name string) (string, error) {
	if name == "Pepper" {
		return "20", nil
	}

	if name == "Floyd" {
		return "10", nil
	}

	return "", errors.New("no player found")
}
