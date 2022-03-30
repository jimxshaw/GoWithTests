package main

// #1

import "fmt"

const greetingPrefix = "Hello "

func Hello(name string, language string) string {
	if name == "" {
		name = "Go"
	}

	if language == "Spanish" {
		return "Hola " + name
	}
	return greetingPrefix + name
}

func main() {
	fmt.Println(Hello("James", "English"))
}
