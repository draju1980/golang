package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#ff12G",
		"blue":  "#ff12B",
	}

	printMap(colors)
}
func printMap(c map[string]string) {
	for key, value := range c {
		fmt.Println("value for", key, "is", value)
	}
}
