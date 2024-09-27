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

var errSingleChar = errors.New("More than one input, please enter a single character: (a, s, m, etc)")
var errNoChar error = errors.New("No input, please enter a single character: (a, s, m, etc)")
var errFailedRead error = errors.New("Failed to read input, Shutting down")

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

func generateMenu() {
	fmt.Println("Pick an operation:")
	fmt.Println("a) Add")
	fmt.Println("s) Subtract")
	fmt.Println("m) Multiply")
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
			}
		}
	}

}
