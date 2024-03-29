package sel

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}

}

func ping(url string) chan struct{} {
	c := make(chan struct{})
	go func() {
		defer close(c)
		http.Get(url)
	}()
	return c
}
