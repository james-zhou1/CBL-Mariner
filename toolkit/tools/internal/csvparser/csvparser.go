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
func ParseAndExport(filename string) {
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

	// create_worker_chroot_path := "/home/james/repos/CBL-Mariner/toolkit/tools/internal/timestamp/results/create_worker_chroot.csv"
	image_config_validator_path := "/home/james/repos/CBL-Mariner/toolkit/tools/internal/timestamp/results/imageconfigvalidator.csv"
	image_pkg_fetcher_path := "/home/james/repos/CBL-Mariner/toolkit/tools/internal/timestamp/results/imagepkgfetcher.csv"
	imager_path := "/home/james/repos/CBL-Mariner/toolkit/tools/internal/timestamp/results/imager.csv"
	roast_path := "/home/james/repos/CBL-Mariner/toolkit/tools/internal/timestamp/results/roast.csv"
	// ParseAndExport(create_worker_chroot_path)
	ParseAndExport(image_config_validator_path)
	ParseAndExport(image_pkg_fetcher_path)
	ParseAndExport(imager_path)
	ParseAndExport(roast_path)

	fmt.Printf("%v\n", timeArray)

	startTime, err := time.Parse(time.RFC1123, "Fri, 01 Jul 2022 17:56:44 EDT")
	if err != nil {
		panic(err)
	}
	endTime, err := time.Parse(time.RFC1123, "Fri, 01 Jul 2022 18:01:44 EDT")
	if err != nil {
		panic(err)
	}
	difference := endTime.Sub(startTime)
	fmt.Println(difference)
}
