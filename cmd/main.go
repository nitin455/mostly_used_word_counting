package main

import (
	"fmt"
	"mostly_used_word_counting/pkg"
)

func main() {
	fmt.Println("Welcome to word calculating utility!!!")

	for {
		fmt.Println("Press Y to Enter New Word")
		fmt.Println("Press N to Exit")
		var text string
		var breakLoop bool

		// Taking input from user
		fmt.Scanln(&text)
		switch text {
		case "Y":
			fmt.Println("Enter Number Of Words You Want To Enter.")
			var count int
			fmt.Scanf("%d", &count)
			var words []string
			for i := 0; i < count+1; i++ {
				if i != 0 {
					fmt.Printf("Enter %d Word.\n", i)
				}

				var input string
				// Taking input from user
				fmt.Scanln(&input)
				words = append(words, input)
			}
			pkg.CountMostlyUsedWords(words)
		case "N":
			breakLoop = true
		default:
			fmt.Println("Invalid input")
		}

		if breakLoop {
			break
		}
	}
}
