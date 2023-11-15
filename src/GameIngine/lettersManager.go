package GameIngine

import (
	"fmt"
	"strings"
)

var guessedLetters []rune

func ToLowerCase(word []rune) []rune {
	for i := 0; i < len(word); i++ {
		word[i] = rune(strings.ToLower(string(word[i]))[0])
	}
	return word
}

func LetterChoice() []rune {
	var letter string

	for {
		fmt.Scan(&letter)

		if len(letter) != 1 {
			fmt.Println("Attention ! Tu dois mettre qu'une seule lettre")
			continue
		}

		letterRune := []rune(letter)[0]

		if contains(guessedLetters, letterRune) {
			fmt.Println("Attention ! Tu as déjà choisi cette lettre. Choisis une nouvelle lettre.")
		} else {
			guessedLetters = append(guessedLetters, letterRune)
			break
		}
	}

	return []rune(letter)
}

func contains(arr []rune, x rune) bool {
	for _, a := range arr {
		if a == x {
			return true
		}
	}
	return false
}
