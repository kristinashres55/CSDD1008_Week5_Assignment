package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

		// Placeholder for unit and conversion
		unit := parts[1]
		fmt.Println("Temperature:", temp, "Unit:", unit)
	}
}
