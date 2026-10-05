package main

import "fmt"

const en = "Hello, "
const es = "Hola, "
const pt = "Olá, "

func prefixGrettings(language string) (prefix string) {
	switch language {
	case "es":
		prefix = es
	case "pt":
		prefix = pt
	default:
		prefix = en
	}
	return
}

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return prefixGrettings(language) + name + "."
}

func main() {
	fmt.Println(Hello("Matheus", "es"))
}
