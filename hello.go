package main

// #1 Hello, World

import "fmt"

const langSpanish = "Spanish"
const greetingEnglishPrefix = "Hello "
const greetingSpanishPrefix = "Hola "

func Hello(name string, language string) string {
	if name == "" {
		name = "Go"
	}

	if language == langSpanish {
		return greetingSpanishPrefix + name
	}

	return greetingEnglishPrefix + name
}

func main() {
	fmt.Println(Hello("James", "English"))
}
