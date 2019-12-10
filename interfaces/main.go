package main

import "fmt"

type englishBot struct {
}
type spanishBot struct {
}

type bot interface {
	getGreeting() string
	printd()
}

func (e spanishBot) printd() {
	fmt.Println("spannish bot running")
}

func (e englishBot) getGreeting() string {
	return "hello there"
}
func (e spanishBot) getGreeting() string {
	return "Hola!"
}
func findPrintGreeting(b bot) {
	b.printd()
}
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func main() {
	sb := spanishBot{}
	// eb := english/Bot{}
	printGreeting(sb)
	findPrintGreeting(sb)
}
