package main

// #1 Hello, World

import "fmt"

const langSpanish = "Spanish"
const langFrench = "French"
const greetingEnglishPrefix = "Hello "
const greetingSpanishPrefix = "Hola "
const greetingFrenchPrefix = "Bonjour "

func Hello(name string, language string) string {
	if name == "" {
		name = "Go"
	}

	prefix := greetingEnglishPrefix

	switch language {
	case langSpanish:
		prefix = greetingSpanishPrefix
	case langFrench:
		prefix = greetingFrenchPrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("James", "English"))
}
