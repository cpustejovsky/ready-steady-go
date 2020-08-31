package main

import (
	"os"
	"time"

	"github.com/cpustejovsky/ready-steady-go/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
