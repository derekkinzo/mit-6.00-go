package main

import (
	"fmt"
	"strings"
)

func main() {
	anLetters := "aefhilmnorsxAEFHILMNORSX"

	var word string // The word entered
	var times int

	fmt.Print("I will cheer for you! Enter a word: ")
	fmt.Scanf("%s", &word)

	fmt.Print("Enthusiasm level (1-10): ")
	fmt.Scanf("%d", &times)

	for _, rune := range word {
		if strings.Contains(anLetters, string(rune)) {
			fmt.Printf("Give me an %c ! %c\n", rune, rune)
		} else {
			fmt.Printf("Give me a %c ! %c\n", rune, rune)
		}
	}

	fmt.Println("What does that spell?")

	for i := 0; i < times; i++ {
		println(word, "!!!")
	}
}
