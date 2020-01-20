package main

import "fmt"

const helloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const spanish = "Spanish"

// Hello func return "Hello, World"
func Hello(name string, language string) string  {
	if name == "" {
		name = "World"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}
	return helloPrefix + name
}

func main()  {
	fmt.Println(Hello("Hello", ""))
}