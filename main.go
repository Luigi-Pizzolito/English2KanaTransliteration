package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Create an instance of AllToKana
	allToKana := NewAllToKana()

	// Listen to stdin indefinitely
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break // Exit loop on error
		}

		// Call convertString function with the accumulated line
		result := allToKana.Convert(line)

		// Output the result
		fmt.Print(result+"\n")
	}
}