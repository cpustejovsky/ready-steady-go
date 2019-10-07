package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	//can't use dot syntax because map's keys are typed!
	colors["z-goblue"] = "#5dc9e2"
	fmt.Println(colors)
	var zeroValMap map[string]string
	fmt.Println(zeroValMap)
	newColors := make(map[int]string)
	newColors[1] = "#5dc9e2"
	fmt.Println(newColors)
	delete(colors, "z-goblue")
	fmt.Println(colors)
}
