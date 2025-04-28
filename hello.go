package main

import "fmt"

const (
	spanish = "Spanish"
	french = "French"
	swedish = "Swedish"

	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix = "Bonjour, "
	swedishHelloPrefix = "Hej, "
)

func Hello(name, lang string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(lang) + name
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	case swedish:
		prefix = swedishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}

	return

}

func main() {
	fmt.Println(Hello("Chris", "English"))
}