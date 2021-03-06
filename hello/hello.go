package main

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

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case langSpanish:
		prefix = greetingSpanishPrefix
	case langFrench:
		prefix = greetingFrenchPrefix
	case langChinese:
		prefix = greetingChinesePrefix
	default:
		prefix = greetingEnglishPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("James", "English"))
}
