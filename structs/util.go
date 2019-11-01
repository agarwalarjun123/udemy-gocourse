package main

import "fmt"

type contactInfo struct {
	email string
	zip   string
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func update(p *person, name string) {
	(*p).firstName = name
}
func (p person) printPerson() {
	fmt.Printf("%+v", p)
}
func (p *person) updateName(f string) {
	p.firstName = f
}
func a(s []string) int {
	return len(s)
}
