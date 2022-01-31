package why

import "fmt"

var greetings = map[string]string{
	"es": "Hola",
	"fr": "Bonjour",
}

func greeting(l string) string {
	greeting, f := greetings[l]
	if !f {
		return "Hello"
	}
	return greeting
}

func Hello(name, language string) string {
	return fmt.Sprintf("%s, %s", greeting(language), name)
}
