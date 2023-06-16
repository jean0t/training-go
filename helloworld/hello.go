package main

import "fmt"

const (
    japanese = "Japanese"
    french = "French"
    spanish = "Spanish"
  japaneseHelloPrefix = "こんにちは"
  englishHelloPrefix = "Hello, "
  frenchHelloPrefix = "Bonjour, "
  spanishHelloPrefix = "Hola, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language{
		case french:
			prefix = frenchHelloPrefix
		case spanish:
			prefix = spanishHelloPrefix
        case japanese:
            prefix = japaneseHelloPrefix

		default:
			prefix = englishHelloPrefix
	}
	return
}


func main() {
	fmt.Println(Hello("Jeanot", "Japanese"))
}
