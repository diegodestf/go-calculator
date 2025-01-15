package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero is not allowed")
	}
	return a / b, nil
}

func validateOperators(operator string, validOperators []string) bool {
	for _, validOp := range validOperators {
		if operator == validOp {
			return true
		}
	}
	return false
}

func calculate() {
	var operator string
	var num1, num2 float64

	reader := bufio.NewReader(os.Stdin)

	validOperators := []string{"+", "-", "*", "/"}

	for {
		fmt.Println("Please enter the an operator: +, -, *, /. Or type 'exit' to quit.")
		fmt.Scan(&operator)
		operator = strings.ToLower(operator)

		if operator == "exit" {
			fmt.Println("Goodbye!")
			break
		}

		if !validateOperators(operator, validOperators) {
			fmt.Println("Oops! Looks like you entered an invalid operator. Please try again")
			continue
		}

		fmt.Println("Please enter the first number:")
		_, err := fmt.Scan(&num1)
		if err != nil {
			fmt.Println("Oops! Looks like you entered an invalid number. Please try again.")
			reader.ReadString('\n')
			continue
		}

		fmt.Println("Please enter the second number:")
		_, err = fmt.Scan(&num2)
		if err != nil {
			fmt.Println("Oops! Looks like you entered an invalid number. Please try again.")
			reader.ReadString('\n')
			continue
		}
		switch operator {
		case "+":
			fmt.Printf("The result is %.2f\n", add(num1, num2))
		case "-":
			fmt.Printf("The result is %.2f\n", subtract(num1, num2))
		case "*":
			fmt.Printf("The result is %.2f\n", multiply(num1, num2))
		case "/":
			result, err := divide(num1, num2)
			if err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Printf("The result is %.2f\n", result)
			}
		}
	}
}

func main() {
	calculate()
}
