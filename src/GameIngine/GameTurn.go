package GameIngine

import (
	"fmt"
)

func GameLoop(wordToFind []rune) {
	var answer []rune
	var hangStep int

	for loop := 0; loop < len(wordToFind); loop++ {
		answer = append(answer, rune(95))
	}
	Display(answer, hangStep)

	for !IsWordIsGood(wordToFind, answer) && hangStep != 7 {
		WordContainLetter(wordToFind, LetterChoice(), answer, &hangStep)
	}
}

func WordContainLetter(word []rune, letter []rune, answer []rune, hangStep *int) {
	word = ToLowerCase(word)
	letter = ToLowerCase(letter)
	contain := false
	for i := len(word) - 1; i >= 0; i-- {
		if word[i] == letter[0] {
			answer[i] = letter[0]
			contain = true
		}
	}
	if contain {
		fmt.Println("Tu as trouvé une lettre")
	} else {
		fmt.Println("La lettre ne fait pas parti du mot, choisis une nouvelle lettre")
		*hangStep++
	}
	Display(answer, *hangStep)
}

func IsWordIsGood(word []rune, answer []rune) bool {
	word = ToLowerCase(word)
	answer = ToLowerCase(answer)
	for i := len(word) - 1; i > 0; i-- {
		if word[i] != answer[i] {
			return false
		}
	}
	fmt.Println("Bien joué ! Tu as trouvé le bon mot qui était:", string(word))
	return true
}
