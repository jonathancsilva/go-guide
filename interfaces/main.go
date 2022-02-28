package main

import "fmt"

type bot interface {
	getGreenting() string
}

type englishBot struct{}
type spanishBot struct{}

func (englishBot) getGreenting() string {
	// very custom logic for generating an english greeting
	return "Hi there!"
}

func (spanishBot) getGreenting() string {
	return "Hola!"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreenting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
