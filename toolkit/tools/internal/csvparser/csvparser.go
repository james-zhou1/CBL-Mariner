package csvparser

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

var timeArray [][]string

// Reads a CSV file, and returns data to the terminal
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

func ParseCSV() {
	wd, _ := os.Getwd()
	idx := strings.Index(wd, "CBL-Mariner/toolkit") // 19 chars
	wd = wd[0 : idx + 19]
	wd += "/tools/internal/timestamp/results/"
	fmt.Printf("%s\n", wd)

	// create_worker_chroot_path := "/home/james/repos/CBL-Mariner/toolkit/tools/internal/timestamp/results/create_worker_chroot.csv"
	image_config_validator_path := "toolkit/tools/internal/timestamp/results/imageconfigvalidator.csv"
	image_pkg_fetcher_path := "toolkit/tools/internal/timestamp/results/imagepkgfetcher.csv"
	imager_path := "toolkit/tools/internal/timestamp/results/imager.csv"
	roast_path := "toolkit/tools/internal/timestamp/results/roast.csv"
	// CSVToArray(create_worker_chroot_path)
	CSVToArray(image_config_validator_path)
	CSVToArray(image_pkg_fetcher_path)
	CSVToArray(imager_path)
	CSVToArray(roast_path)

	// Get the start time from the first timestamp entry
	startTime, err := time.Parse(time.UnixDate, timeArray[0][4])
	if err != nil {
		panic(err)
	}

	// Get the end time from the last timestamp entry
	endTime, err := time.Parse(time.UnixDate, timeArray[len(timeArray)-1][5])
	if err != nil {
		panic(err)
	}

	// Get the time difference (total build time)
	difference := endTime.Sub(startTime)

	// Print timestamps
	for i := 0; i < len(timeArray); i++ {
		fmt.Println(timeArray[i][0] + " " + timeArray[i][1] + " took " + timeArray[i][3] + ". ")
	}

	fmt.Println("The full build duration was " + difference.String() + ".")
}
