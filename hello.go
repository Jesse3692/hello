package main

import "fmt"

const englishHelloPrefix = "Hello, "

// Hello func return "Hello, World"
func Hello(name string) string  {
	if name == "" {
		return "Hello, World"
	}
	return englishHelloPrefix + name
}

func main()  {
	fmt.Println(Hello("Hello"))
}