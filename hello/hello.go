package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const englishWorld = "World"
const frenchWorld = "Monde"
const spanishWorld = "Mundo"
const spanish = "Spanish"
const french = "French"

func Hello(name, language string) string {
	return greetingPrefix(language) + worldWord(name, language) + "!"
}

func worldWord(name string, language string) (worldname string) {
	if name == "" {
		switch language {
		case spanish:
			worldname = spanishWorld
		case french:
			worldname = frenchWorld
		default:
			worldname = englishWorld
		}
	} else {
		worldname = name
	}
	return
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}