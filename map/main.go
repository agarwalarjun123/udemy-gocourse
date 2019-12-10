package main

import "fmt"

func main() {
	colors :=
		map[string]string{
			"red":   "ff0000",
			"green": "00ff00",
			"blue":  "0000ff",
		}
	c := make(chan bool)
	go Modify(colors, c)
	go printMap(colors, c)
	<-c
	<-c
}
func printMap(c map[string]string, d chan bool) {
	for color, hex := range c {
		fmt.Printf("%s -> %s\n", color, hex)
	}
	c["blue"] = "001100"
	d <- true
}

func Modify(c map[string]string, d chan bool) {

	c["red"] = "fff111"
	fmt.Print("Modify Function\n")
	for color, hex := range c {
		fmt.Printf("1%s -> %s\n", color, hex)
	}
	d <- true
}
