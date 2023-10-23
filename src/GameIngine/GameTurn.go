package GameIngine

import (
	"fmt"
)

func GameLoop(wordToFind []rune) {
	var answer []string
	var letter string
	var hangStep int
	//var letterNotWorking []string
	for loop := 0; loop < len(wordToFind); loop++ {
		answer = append(answer, "_")
	}
	fmt.Println(answer)
	for rune(answer[0][0]) != wordToFind[0] || hangStep != 7 {
		fmt.Scan(&letter)
		if len(letter) > 1 {
			fmt.Println("Pas bon")
			fmt.Scan(&letter)
		}
		//if strings.Contains(wordToFind, letter) {

		//} else {
		//	letterNotWorking = append(letterNotWorking, letter)
		//}
	}
}
