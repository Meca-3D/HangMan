package main

import (
	"hangman/GameIngine"
	"hangman/wordSelector"
)

func main() {
	var wordlist []string
	wordlist = wordSelector.ReadFile()
	word := wordSelector.ChoseRandomWord(wordlist)
	GameIngine.GameLoop(word)
}
