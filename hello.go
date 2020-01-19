package main

import "fmt"

// Hello func return "Hello, World"
func Hello(name string) string  {
	return "Hello, " + name
}

func main()  {
	fmt.Println(Hello("Hello"))
}