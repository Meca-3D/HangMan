package GameIngine

import "fmt"

func Display(runeToDisplay []rune, step int) {
	fmt.Println("Trouve le mot ci-dessous: ")

	var word string
	for i := len(runeToDisplay) - 1; i >= 0; i-- {
		word = string(runeToDisplay[i]) + " " + word
	}
	fmt.Println(word, "\n")

	fmt.Println("Liste des lettres utilisées:\n", string(guessedLetters))

	switch step {
	case 0:
		fmt.Println("  ┌───────┐		")
		fmt.Println("  │               	")
		fmt.Println("  │               	")
		fmt.Println("  │               	")
		fmt.Println("  │               	")
		fmt.Println("  │               	")
		fmt.Println("  └───────┘\n		")
	case 1:
		fmt.Println("  ┌───────┐		")
		fmt.Println("  │/       		")
		fmt.Println("  │               	")
		fmt.Println("  │               	")
		fmt.Println("  │               	")
		fmt.Println("  │               	")
		fmt.Println("  └───────┘\n		")
	case 2:
		fmt.Println("  ┌───────┐		")
		fmt.Println("  │/      O		")
		fmt.Println("  │               	")
		fmt.Println("  │               	")
		fmt.Println("  │               	")
		fmt.Println("  │               	")
		fmt.Println("  └───────┘\n		")
	case 3:
		fmt.Println("  ┌───────┐		")
		fmt.Println("  │/      O		")
		fmt.Println("  │       |		")
		fmt.Println("  │               	")
		fmt.Println("  │               	")
		fmt.Println("  │               	")
		fmt.Println("  └───────┘\n		")
	case 4:
		fmt.Println("  ┌───────┐		")
		fmt.Println("  │/      O		")
		fmt.Println("  │      /|		")
		fmt.Println("  │               	")
		fmt.Println("  │               	")
		fmt.Println("  │               	")
		fmt.Println("  └───────┘\n		")
	case 5:
		fmt.Println("  ┌───────┐")
		fmt.Println("  │/      O")
		fmt.Println("  │      /|\\")
		fmt.Println("  │               	")
		fmt.Println("  │               	")
		fmt.Println("  │               	")
		fmt.Println("  └───────┘\n		")
	case 6:
		fmt.Println("  ┌───────┐		")
		fmt.Println("  │/      O		")
		fmt.Println("  │      /|\\		")
		fmt.Println("  │      / 		")
		fmt.Println("  │               	")
		fmt.Println("  │               	")
		fmt.Println("  └───────┘\n		")
	default:
		fmt.Println("  ┌───────┐		")
		fmt.Println("  │/      O		")
		fmt.Println("  │      /|\\		")
		fmt.Println("  │      / \\		")
		fmt.Println("  │               	")
		fmt.Println("  │               	")
		fmt.Println("  └───────┘\n		")
		fmt.Println("	GAME OVER		")
	}
}
