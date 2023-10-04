package main

import (
	"fmt"
	"hangman/GameIngine"
	"hangman/wordSelector"
)

func main() {
	var wordlist []string
	wordlist = wordSelector.ReadFile()
	word := wordSelector.ChoseRandomWord(wordlist)
	fmt.Println(word)
	GameIngine.GameLoop(word)
}
