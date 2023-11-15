package main

import (
	"fmt"
	"hangman/GameIngine"
	"hangman/wordSelector"
)

func main() {

	fmt.Println(" _________________________________________________________________")
	fmt.Println("    _    _   ")
	fmt.Println("   | |  | |")
	fmt.Println("   | |__| |   __ _   _ __     __ _   _ __ ___     __ _   _ __")
	fmt.Println("   |  __  |  / _` | | '_ \\   / _` | | '_ ` _ \\   / _` | | '_ \\")
	fmt.Println("   | |  | | | (_| | | | | | | (_| | | | | | | | | (_| | | | | |")
	fmt.Println("   |_|  |_|  \\__,_| |_| |_|  \\__, | |_| |_| |_|  \\__,_| |_| |_|")
	fmt.Println("                              __/ |")
	fmt.Println(" _________________________________________________________________ \n")

	fmt.Println("Choisissez le mode de jeu:")
	fmt.Println("1. Solo")
	fmt.Println("2. Duo")

	var choise int
	fmt.Scan(&choise)

	switch choise {
	case 1:
		var wordlist []string
		wordlist = wordSelector.ReadFile()
		word := wordSelector.ChoseRandomWord(wordlist)
		GameIngine.GameLoop(word)
	case 2:
		fmt.Println("Le joueur 1 doit choisir le mot à faire deviner au joueur 2:")
		wordToGuess := wordSelector.ChooseWord()
		GameIngine.GameLoop([]rune(wordToGuess))
	default:
		fmt.Println("Par défaut le mode de jeu est solo.")
		var wordlist []string
		wordlist = wordSelector.ReadFile()
		word := wordSelector.ChoseRandomWord(wordlist)
		GameIngine.GameLoop(word)
	}
}
