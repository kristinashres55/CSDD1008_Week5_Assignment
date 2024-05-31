package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CelsiusToFahrenheit converts a temperature from Celsius to Fahrenheit.
func CelsiusToFahrenheit(c float64) float64 {
	return c*9/5 + 32

}

// FahrenheitToCelsius converts a temperature from Fahrenheit to Celsius.
func FahrenheitToCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
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

		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Println("Invalid input. Please enter the temperature followed by its unit (C or F).")
			continue
		}

		temp, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			fmt.Println("Invalid temperature. Please enter a valid numeric temperature.")
			continue
		}

		unit := parts[1]
		switch strings.ToUpper(unit) {
		case "F":
			celsius := FahrenheitToCelsius(temp)
			fmt.Printf("%.2f C\n", celsius)
		default:
			fmt.Println("Invalid unit. Please specify 'C' for Celsius or 'F' for Fahrenheit.")
		}
	}
}
