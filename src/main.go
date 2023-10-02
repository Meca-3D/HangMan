package main

import (
	"fmt"
	"hangman/wordSelector"
)

func main() {
	var wordlist []string
	wordlist = wordSelector.ReadFile()
	fmt.Println(wordSelector.ChoseRandomWord(wordlist))
}
