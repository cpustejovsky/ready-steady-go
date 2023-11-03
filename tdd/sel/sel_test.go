package sel_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cpustejovsky/ready-steady-go/tdd/sel"
)

func TestRacer(t *testing.T) {

	slowServer := makeDelayedServer(60 * time.Millisecond)
	fastServer := makeDelayedServer(50 * time.Millisecond)

	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got, err := sel.Racer(slowURL, fastURL, 10*time.Second)
	if err != nil {
		t.Error(err)
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
