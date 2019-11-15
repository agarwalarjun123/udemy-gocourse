package main

import "fmt"

type englishBot struct {
}
type spanishBot struct {
}

type bot interface {
	getGreeting() string
}
type r struct {
	f bot
}

func (e r) getGreeting() string {
	return "a"
}

func (e englishBot) getGreeting() string {
	return "hello there"
}
func (e spanishBot) getGreeting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	sb := spanishBot{}
	eb := englishBot{}
	printGreeting(eb)
	printGreeting(sb)
}
