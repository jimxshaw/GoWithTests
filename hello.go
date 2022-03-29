package main

// #1

import "fmt"

const greetingPrefix = "Hello "

func Hello(name string) string {
	if name == "" {
		name = "Go"
	}
	return greetingPrefix + name
}

func main() {
	fmt.Println(Hello("James"))
}
