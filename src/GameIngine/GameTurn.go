package GameIngine

import (
	"fmt"
)

func GameLoop(wordToFind []rune) {
	var answer []rune
	var hangStep int
	//var letterNotWorking []string
	for loop := 0; loop < len(wordToFind); loop++ {
		answer = append(answer, rune(95))
	}
	Display(answer, hangStep)
	for IsWordIsGood(wordToFind, answer) || hangStep != 7 {
		WordContainLetter(wordToFind, LetterChoice(), answer)
	}
}

func LetterChoice() []rune {
	var letter string
	fmt.Scan(&letter)
	for len(letter) > 1 || len(letter) <= 0 {
		fmt.Println("Pas bon")
		fmt.Scan(&letter)
	}
	returnLetter := []rune(letter)
	return returnLetter
}

func WordContainLetter(word []rune, letter []rune, answer []rune) {
	contain := false
	for i := len(word) - 1; i >= 0; i-- {
		if word[i] == letter[0] {
			answer[i] = letter[0]
			contain = true
			Display(answer, 5)
		}
	}
	if contain == true {
		fmt.Println("tu as trouver une lettre")
	}
}

func Display(runeToDisplay []rune, step int) {
	var word string
	for i := len(runeToDisplay) - 1; i >= 0; i-- {
		word = string(runeToDisplay[i]) + " " + word
	}
	fmt.Println(step)
	fmt.Println(word)
}

func IsWordIsGood(word []rune, answer []rune) bool {
	for i := len(word) - 1; i > 0; i-- {
		fmt.Println(i)
		if word[i] != answer[i] {
			return false
		}
	}
	fmt.Println("Mot trouver")
	return true
}
