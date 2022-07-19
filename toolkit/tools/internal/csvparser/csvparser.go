package csvparser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	timeArray [][]string
	files     = []string{"/imageconfigvalidator.csv", "/imagepkgfetcher.csv", "/imager.csv", "/roast.csv"}
)

// Reads a CSV file and appends line by line to array
func CSVToArray(filename string) {
	file, err := os.Open(filename)

	if err != nil {
		fmt.Println("failed to open csv file")
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		timeArray = append(timeArray, strings.Split(fileScanner.Text(), ","))
	}

}

// Return list of file paths
func FilepathsToArray(parentDir string) []string {
	// wd, _ := os.Getwd()
	// idx := strings.Index(wd, "CBL-Mariner/toolkit") // 19 chars
	// wd = wd[0 : idx+19]
	// wd += "/tools/internal/timestamp/results/"

	image_config_validator_path := parentDir + "/imageconfigvalidator.csv"
	image_pkg_fetcher_path := parentDir + "/imagepkgfetcher.csv"
	imager_path := parentDir + "/imager.csv"
	roast_path := parentDir + "/roast.csv"

	fileArray := []string{image_config_validator_path, image_pkg_fetcher_path, imager_path, roast_path}

	return fileArray
}

// Take list of file paths, parse, and output log to terminal
func OutputCSVLog(parentDir string) {
	var startTime time.Time
	init_file, err := os.Stat(parentDir + "/init")

	// Format each file to array format
	for _, file := range files {
		CSVToArray(parentDir + file)
	}

	// Get the start and end time from the first timestamp entry
	if os.IsNotExist(err) {
		fmt.Printf("start: %s\n", timeArray[0][4])
		startTime, err = time.Parse(time.UnixDate, timeArray[0][4])
	} else {
		fmt.Printf("start: %s\n", init_file.ModTime().Format(time.UnixDate))
		startTime = init_file.ModTime()
	}
	fmt.Printf("end: %s\n", timeArray[len(timeArray)-1][5])

	if err != nil {
		panic(err)
	}

	endTime, err := time.Parse(time.UnixDate, timeArray[len(timeArray)-1][5])
	if err != nil {
		panic(err)
	}

	// Get the time difference (aka. total build time)
	difference := endTime.Sub(startTime)

	// Print timestamps
	for i := 0; i < len(timeArray); i++ {
		fmt.Println(timeArray[i][0] + " " + timeArray[i][1] + " took " + timeArray[i][3] + ". ")
	}

	fmt.Println("The full build duration was " + difference.String() + ".")
}
