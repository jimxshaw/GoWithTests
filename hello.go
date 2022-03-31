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

	if language == langSpanish {
		return greetingSpanishPrefix + name
	}

	if language == langFrench {
		return greetingFrenchPrefix + name
	}

	return greetingEnglishPrefix + name
}

func main() {
	fmt.Println(Hello("James", "English"))
}
