package GameIngine

import (
	"fmt"
	"strings"
)

func GameLoop(wordToFind []string) {
	var answer []string
	var letter string
	var hangStep int
	var letterNotWorking []string
	loop := 0
	for loop < len(wordToFind) {
		loop += 1
		answer = append(answer, "_")
	}
	fmt.Println()
	for answer[0] != wordToFind[0] || hangStep != 7 {
		fmt.Scan(&letter)
		if strings.Contains(wordToFind[0], letter) {

		} else {
			letterNotWorking = append(letterNotWorking, letter)
		}
	}
}
