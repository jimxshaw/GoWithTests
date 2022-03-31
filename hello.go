package main

// #1 Hello, World

import "fmt"

const langSpanish = "Spanish"
const langFrench = "French"
const langChinese = "Chinese"
const greetingEnglishPrefix = "Hello "
const greetingSpanishPrefix = "Hola "
const greetingFrenchPrefix = "Bonjour "
const greetingChinesePrefix = "Ni hao "

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
	case langChinese:
		prefix = greetingChinesePrefix
	}

	return prefix + name
}

func main() {
	fmt.Println(Hello("James", "English"))
}
