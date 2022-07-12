package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"strings"
	"time"

	"github.com/gosuri/uiprogress"
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
	uiprogress.Start() // start rendering
	bar := uiprogress.AddBar(22)
	bar.AppendCompleted()
	bar.PrependElapsed()

	wd, _ := os.Getwd()
	idx := strings.Index(wd, "CBL-Mariner")
	wd = wd[0 : idx+11]
	targetDir = wd + "/build/timestamp/"

	// Use an infinite for loop to watch out for new updates
	for {
		// run this iteration periodically for smaller overhead
		time.Sleep(1 * time.Second)

		// Check if the target directory exists. Assume we only need to check one directory for now.
		_, err := os.Stat(targetDir)
		if os.IsNotExist(err) {
			continue
		}

		// Check update for each target file.
		for filePath, _ := range targetCSV {
			// currStat (the deleted 1st variable) will be important when we try to print out info in front of the progress bars.
			_, err := os.Stat(targetDir + filePath) 
			// Check if the file exists.
			if os.IsNotExist(err) {
				continue
			}
			// getUpdate(currStat, idx, filePath)
			currNumLines := getNumLines(targetDir + filePath)
			if currNumLines != targetCSV[filePath][0] {
				targetCSV[filePath][0] = currNumLines
				// fmt.Printf("%s has %d lines \n", currStat.Name(), currNumLines)
				bar.Incr()
			}
			// getUpdate(currStat, idx, filePath)
		}
	}

}

// Check if the file has been updated, and get updated contents if it did.
// Assumption: the file and its parent directories of the file have been created.
func getUpdate(currStat fs.FileInfo, idx int, filePath string) {
	currNumLines := getNumLines(targetDir + filePath)
	if currNumLines != targetCSV[filePath][0] {
		targetCSV[filePath][0] = currNumLines
		// fmt.Printf("%s has %d lines \n", currStat.Name(), currNumLines)
	}
}

// Naive implementation (potentially inefficient for larger files)
func getNumLines(filepath string) int64 {
	file, _ := os.Open(filepath)
	fileScanner := bufio.NewScanner(file)
	count := 0

	for fileScanner.Scan() {
		count++
	}

	return int64(count)
}
