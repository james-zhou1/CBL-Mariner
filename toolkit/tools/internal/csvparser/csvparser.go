package csvparser

import (
	"bufio"
	"fmt"
	"os"
)

// Reads a CSV file, and returns data to the terminal
func ParseAndExport() {
	file, err := os.Open("file_data.csv")

	if err != nil {
		fmt.Println("failed to open csv file")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		fmt.Println((fileScanner.Text()))
	}

}
