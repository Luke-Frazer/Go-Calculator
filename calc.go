package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var errSingleChar = errors.New("more than one input, please enter a single character: (a, s, m, etc)")
var errNoChar error = errors.New("no input, please enter a single character: (a, s, m, etc)")
var errFailedRead error = errors.New("failed to read input, shutting down")

func getInput(reader io.Reader) (string, error) {
	scanner := bufio.NewScanner(reader)
	if scanner.Scan() {
		var fullInputLine string = scanner.Text()
		var inputs []string = strings.Fields(fullInputLine)
		if len(inputs) < 1 {
			return "", errNoChar
		}
		if len(inputs) > 1 {
			return "", errSingleChar
		}
		return inputs[0], nil
	}
	return "", errFailedRead
}

func getInputs(reader io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(reader)
	if scanner.Scan() {
		var fullInputLine string = scanner.Text()
		var inputs []string = strings.Fields(fullInputLine)
		if len(inputs) < 1 {
			return nil, errNoChar
		}
		return inputs, nil
	}
	return nil, errFailedRead
}

func convertToNumbers(inputs ...string) ([]int, error) {
	numbers := []int{}
	for _, input := range inputs {
		convertedInput, err := strconv.Atoi(input)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, convertedInput)
	}
	return numbers, nil
}

func addValues(numbers ...int) int {
	acc := 0
	for _, number := range numbers {
		acc += number
	}
	return acc
}

func multValues(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	}
	acc := 1
	for _, number := range numbers {
		acc *= number
	}
	return acc
}

func subValues(numbers ...int) int {
	if len(numbers) == 0 {
		return 0
	} else if len(numbers) == 1 {
		return numbers[0]
	}
	acc := numbers[0]
	for i := 1; i < len(numbers); i++ {
		acc -= numbers[i]
	}
	return acc
}

func generateMenu() {
	fmt.Println("Pick an operation:")
	fmt.Println("a) Add")
	fmt.Println("s) Subtract")
	fmt.Println("m) Multiply")
	fmt.Println("q) Quit")
}

func main() {
	continueInput := true
	for continueInput {
		generateMenu()
		scannedItems, err := getInput(os.Stdin)
		if err == errNoChar || err == errSingleChar {
			fmt.Println(err)
			continue
		} else if err != nil {
			log.Fatal("Error: ", err)
		} else {
			switch scannedItems {
			case "a":
				fmt.Print("Enter numbers to add: ")
				inputs, err := getInputs(os.Stdin)
				if err != nil {
					fmt.Println(err)
					continue
				}
				numbers, err := convertToNumbers(inputs...)
				if err != nil {
					fmt.Println(err)
					continue
				}
				answer := addValues(numbers...)
				fmt.Println("Your answer: ", answer)

			case "s":
				fmt.Print("Enter numbers to subtract: ")
				inputs, err := getInputs(os.Stdin)
				if err != nil {
					fmt.Println(err)
					continue
				}
				numbers, err := convertToNumbers(inputs...)
				if err != nil {
					fmt.Println(err)
					continue
				}
				answer := subValues(numbers...)
				fmt.Println("Your answer: ", answer)
			case "m":
				fmt.Print("Enter numbers to multiply: ")
				inputs, err := getInputs(os.Stdin)
				if err != nil {
					fmt.Println(err)
					continue
				}
				numbers, err := convertToNumbers(inputs...)
				if err != nil {
					fmt.Println(err)
					continue
				}
				answer := multValues(numbers...)
				fmt.Println("Your answer: ", answer)
			case "q":
				fmt.Println("Quitting Program...")
				os.Exit(0)
			}
		}
	}

}
