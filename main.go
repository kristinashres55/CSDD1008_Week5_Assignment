package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// CelsiusToFahrenheit converts a temperature from Celsius to Fahrenheit.
func CelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter a temperature with its unit (e.g., '32 F' or '100 C'), or type 'exit' to quit: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if strings.ToLower(input) == "exit" {
			fmt.Println("Exiting the program.")
			break
		}
	}
}
