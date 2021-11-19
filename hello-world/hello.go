package main

import "fmt"

const (
	english            = "English"
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchhHelloPrefix = "Bonjour, "
)

func Hello(name, lang string) (prefix string) {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case english:
		prefix = englishHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchhHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("David", ""))
}
