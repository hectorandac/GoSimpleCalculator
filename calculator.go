package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func main() {

	fmt.Println("Simple calculator")
	fmt.Println("([1] addition, [2] subtraction, [3] division, [4] multiplication)")
	fmt.Println("---------------------")
	selectedOption := 0
	for selectedOption == 0 {
		switch captureUserInput() {
		case "1":
			selectedOption = 1
			fmt.Println("Selected operation: addition")
		case "2":
			selectedOption = 2
			fmt.Println("Selected operation: subtraction")
		case "3":
			selectedOption = 3
			fmt.Println("Selected operation: division")
		case "4":
			selectedOption = 4
			fmt.Println("Selected operation: multiplication")
		default:
			fmt.Println("Invalid selection try again")
		}
	}

	switch selectedOption {
	case 1:
		fmt.Printf("Result: %f\n", addition(captureUserNumberInput(), captureUserNumberInput()))
	case 2:
		fmt.Printf("Result: %f\n", subtraction(captureUserNumberInput(), captureUserNumberInput()))
	case 3:
		fmt.Printf("Result: %f\n", division(captureUserNumberInput(), captureUserNumberInput()))
	case 4:
		fmt.Printf("Result: %f\n", multiplication(captureUserNumberInput(), captureUserNumberInput()))
	default:
		fmt.Println("Invalid operation")
	}
}

func captureUserNumberInput() float32 {
	f64, _ := strconv.ParseFloat(captureUserInput(), 32)
	return float32(f64)
}

func captureUserInput() string {
	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	return strings.Replace(text, "\n", "", -1)
}

func addition(num1, num2 float32) float32 {
	return num1 + num2
}

func subtraction(num1, num2 float32) float32 {
	return num1 - num2
}

func division(num1, num2 float32) float32 {
	return num1 / num2
}

func multiplication(num1, num2 float32) float32 {
	return num1 * num2
}
