package main

import "fmt"

func main() {
	// var colors map[string]string
	// colors := make(map[string]string)
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"blue":  "#0000ff",
	}
	printMap(colors)
} 

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Println("Hex code for", k, "is", v)
	}
}
