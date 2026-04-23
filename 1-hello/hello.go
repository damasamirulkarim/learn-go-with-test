package main

import "fmt"

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	greeting := "Hello, %s"
	switch language {
	case "Indonesian":
		greeting = "Halo, %s"
	case "Spanish":
		greeting = "Hola, %s"
	}

	return fmt.Sprintf(greeting, name)
}

func main() {
	fmt.Println(Hello("World", ""))
}