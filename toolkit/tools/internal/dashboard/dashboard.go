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
)

func main() {
	fmt.Println("Starting dashboard")
	uiprogress.Start()
	bar := uiprogress.AddBar(22).AppendCompleted().PrependElapsed()

	wd, _ := os.Getwd()
	idx := strings.Index(wd, "CBL-Mariner")
	wd = wd[0 : idx+11]
	targetDir = wd + "/build/timestamp/"

	for {
		time.Sleep(1 * time.Second)
<<<<<<< HEAD
=======

		// Check if the target directory exists. Assume we only need to check one directory for now.
>>>>>>> refs/remotes/origin/appleli/timestamp
		_, err := os.Stat(targetDir)
		if os.IsNotExist(err) {
			continue
		}
		for filePath, _ := range targetCSV {
<<<<<<< HEAD
			_, err = os.Stat(targetDir + filePath)
=======
			// currStat (the deleted 1st variable) will be important when we try to print out info in front of the progress bars.
			_, err := os.Stat(targetDir + filePath) 
			// Check if the file exists.
>>>>>>> refs/remotes/origin/appleli/timestamp
			if os.IsNotExist(err) {
				continue
			}
			currNumLines := getNumLines(targetDir + filePath)
			if currNumLines != targetCSV[filePath][0] {
				targetCSV[filePath][0] = currNumLines
				// fmt.Printf("%s has %d lines \n", currStat.Name(), currNumLines)
				bar.Incr()
			}
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
