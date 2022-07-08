package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"time"
)

var (
	targetDir string
	targetCSV = map[string][]int64{
		"create_worker_chroot.csv": []int64{0, 4},
		"imageconfigvalidator.csv": []int64{0, 2},
		"imagepkgfetcher.csv":      []int64{0, 9},
		"imager.csv":               []int64{0, 4},
		"roast.csv":                []int64{0, 3},
	}
	// targetCSV = []string{"create_worker_chroot.csv", "imageconfigvalidator.csv", "imagepkgfetcher.csv", "imager.csv", "roast.csv"}
	// CSVSize   = []int64{0, 0, 0, 0, 0}
	// targetJSON = []string{}
)

func main() {
	fmt.Println("Starting dashboard")
	wd, _ := os.Getwd()
	idx := strings.Index(wd, "CBL-Mariner/toolkit")
	wd = wd[0 : idx+19]
	targetDir = wd + "/tools/internal/timestamp/results/"

	// Use an infinite for loop to watch out for new updates
	for {
		// run this iteration periodically for smaller overhead
		time.Sleep(1 * time.Second)

		// Check if the target directory exists. Assume we only need to check one directory for now.
		currStat, err := os.Stat(targetDir)
		if os.IsNotExist(err) {
			// fmt.Printf("The target directory %s doesn't exist. \n", targetDir)
			continue
		}

		// Check update for each target file.
		for filePath, _ := range targetCSV {
			// filePath = targetDir + filePath
			// fmt.Printf("Processing file %s. \n", filePath)
			currStat, err = os.Stat(targetDir + filePath)
			// Check if the file exists.
			if os.IsNotExist(err) {
				// fmt.Printf("File doesn't exist. \n")
				continue
			}
			// getUpdate(currStat, idx, filePath)
			getUpdate(currStat, idx, filePath)
		}
	}

}

// Check if the file has been updated, and get updated contents if it did.
// Assumption: the file and its parent directories of the file have been created.
func getUpdate(currStat fs.FileInfo, idx int, filePath string) {
	currNumLines := getNumLines(filePath)
	if currNumLines != targetCSV[filePath][0] {
		targetCSV[filePath][0] = currNumLines
		
		fmt.Printf("%s has %d lines \n\n", currStat.Name(), currNumLines)
		fmt.Printf("--------\n")
	}
}

// Naive implementation (potentially inefficient for larger files)
func getNumLines(filepath string) int64 {
	file, _ := os.Open(targetDir + filepath)
	fileScanner := bufio.NewScanner(file)
	// fileScanner.Split(bufio.ScanLines) // Tells the scanner to read the file line by line (by default)
	count := 0

	for fileScanner.Scan() {
		count++
		// if there are new lines, print each line
		if count > int(targetCSV[filepath][0]) {
			fmt.Printf("[%d / %d] in %s: %s \n", count, targetCSV[filepath][1], filepath, fileScanner.Text())
		}
	}

	return int64(count)
}
