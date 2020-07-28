package main

import "fmt"

type bot interface {
	getGreeting() string
}

// interfaces are implicit in go
// as long as the concrete type implements
// the interface's methods
type englishBot struct{}
type spanishBot struct{}
type germanBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}
	gb := germanBot{}

	printGreeting(eb)
	printGreeting(sb)
	printGreeting(gb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// fun fact: if ur not using the struct var
// in ur receiver, you can delete it
func (englishBot) getGreeting() string {
	// VERY custom logic for english greeting lol
	return "Hello!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func (germanBot) getGreeting() string {
	return "Hallo!"
}
