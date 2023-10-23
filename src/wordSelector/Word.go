package wordSelector

import (
	"bufio"
	"log"
	"math/rand"
	"os"
)

func ReadFile() []string {
	var word []string
	f, err := os.Open("words.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		word = append(word, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return word
}

func ChoseRandomWord(wordlist []string) []rune {
	var wordInArray string
	randomIndex := rand.Intn(len(wordlist) - 1)
	wordInArray += wordlist[randomIndex]
	s := []rune(wordInArray)
	return s
}
