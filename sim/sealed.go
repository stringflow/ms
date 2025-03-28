package main

import (
	"fmt"
	"math/rand/v2"
	"strconv"
)

type Solution struct {
	left  int
	mid   int
	right int
}

func (s Solution) Compare(o Solution) (matches int) {
	if s.left == o.left {
		matches++
	}

	if s.mid == o.mid {
		matches++
	}

	if s.right == o.right {
		matches++
	}

	return
}

func GenerateSolution() (solution Solution) {
	remaining := 5

	solution.left = rand.IntN(remaining + 1)
	remaining -= solution.left

	solution.mid = rand.IntN(remaining + 1)
	remaining -= solution.mid

	solution.right = remaining
	remaining -= solution.right

	return
}

func GetUserInput() (solution Solution, err error) {
	var input string
	_, err = fmt.Scanln(&input)
	if err != nil {
		return
	}

	input_int, err := strconv.ParseInt(input, 10, 32)
	if err != nil {
		return
	}

	solution.left = int(input_int) / 100
	solution.mid = (int(input_int) % 100) / 10
	solution.right = int(input_int) % 10
	return
}

func main() {
	solution := GenerateSolution()
	attempts := 0

	for attempts < 7 {
		var guess Solution
		var err error
		for {
			fmt.Print("Enter the guess in format 'lmr': ")
			guess, err = GetUserInput()
			if err == nil {
				break
			}
		}

		attempts++
		matches := solution.Compare(guess)

		if matches == 3 {
			fmt.Println("That was the correct combination!")
			return
		} else {
			fmt.Printf("Attempt #%v.\n%v platforms are correct.\n", attempts, matches)
		}
	}

	fmt.Println("You have ran out of guesses.")
}
