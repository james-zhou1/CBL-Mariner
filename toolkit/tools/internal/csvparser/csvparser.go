package csvparser

import (
	"bufio"
	"fmt"
	"os"
)

// Reads a CSV file, and returns data to the terminal
func ParseAndExport(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("failed to open csv file")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		fmt.Println((fileScanner.Text()))
	}

}

func ParseCSV() {

	create_worker_chroot_path := "/home/james/repos/CBL-Mariner/toolkit/tools/internal/timestamp/results/create_worker_chroot.csv"
	ParseAndExport(create_worker_chroot_path)

}
