package main

// #1

import "fmt"

func Hello(name string) string {
	return "Hello " + name + "!"
}

func main() {
	fmt.Println(Hello("James"))
}
