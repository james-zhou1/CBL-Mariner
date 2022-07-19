// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// Tool to track build time during image generation.

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/microsoft/CBL-Mariner/toolkit/tools/internal/exe"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app          = kingpin.New("bldtracker", "A tool that helps track build time of different steps in the makefile.")
	scriptName   = app.Flag("script-name", "The name of the current tool.").Required().String()
	stepName     = app.Flag("step-name", "The name of the current step.").Required().String()
	actionName   = app.Flag("action-name", "The name of the current action.").Default("").String()
	filePath     = app.Flag("file-path", "The folder that stores timestamp csvs.").Required().ExistingDir()	// currently must be absolute
	mode         = app.Flag("mode", "The mode of this tool. Could be 'initialize' ('n') or 'record'('r').").Required().String() 
	completePath string
)

func main() {
	app.Version(exe.ToolkitVersion)
	kingpin.MustParse(app.Parse(os.Args[1:]))

	// Format the script name by removing ".sh".
	var shortName string
        shortName, _, _ = strings.Cut((*scriptName), ".")
	completePath = *filePath + "/" + shortName + ".csv"
	switch *mode {
	case "n":
		initialize()
		break
	case "r":
		record()
		break
	default:
		fmt.Printf("Invalid call. Mode must be 'n' for initialize or 'r' for record. ")
	}
}

// Creates a csv specifically for the shell script mentioned in "scriptName". 
func initialize() {
	file, err := os.Create(completePath)
	if err != nil {
		fmt.Printf("Unable to create file: %s", completePath)
	}
	file.Close()

	// Make a timestamp record right when a shell script starts.
	record()
}

// Records a new timestamp to the specific csv for the specified shell script. 
func record() {
	file, err := os.OpenFile(completePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Unable to open file (may not have been created): %s", completePath)
	}
	defer file.Close()
	
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the timestamp to the csv.
	err = writer.Write([]string{*scriptName, *stepName, *actionName, time.Now().Format(time.UnixDate)})
	if err != nil {
		fmt.Printf("Unable to write to file: %s", completePath)
	}
}
