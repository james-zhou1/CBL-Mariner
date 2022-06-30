// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.

// Tool to track build time during image generation.

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"syscall"
	"time"

	"github.com/microsoft/CBL-Mariner/toolkit/tools/internal/exe"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app        = kingpin.New("bldtracker", "A tool that helps track build time of different steps in the makefile.")
	scriptName = app.Flag("script-name", "The name of the current tool.").Required().String()
	stepName   = app.Flag("step-name", "The name of the current step.").Required().String()
	actionName = app.Flag("action-name", "The name of the current sub-action.").Default("").String()
	filePath   = app.Flag("file-path", "The file that stores timestamp data.").Required().String()                            // currently must be absolute
	mode       = app.Flag("mode", "The mode of this tool. Could be 'initialize' ('n') or 'record'('r').").Required().String() // should I set a default?
)

func main() {
	app.Version(exe.ToolkitVersion)
	kingpin.MustParse(app.Parse(os.Args[1:]))
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

func initialize() {
	mask := syscall.Umask(0)
	defer syscall.Umask(mask)
	err := os.MkdirAll(filepath.Dir(*filePath), 0777)
	if err != nil {
		panic(err)
	}
	file, err := os.Create(*filePath)
	if err != nil {
		fmt.Printf("Unable to create file: %s", *filePath)
	}
	file.Close()
}

func record() {
	file, err := os.OpenFile(*filePath, os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Printf("Unable to open file (may not have been created): %s", *filePath)
	}
	defer file.Close()
	// Create a new csv writer.
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// err = writer.Write([]string{info.toolName, info.stepName, info.actionName, info.duration.String()})
	err = writer.Write([]string{*scriptName, *stepName, *actionName, time.Now().Format(time.RFC1123)})
	if err != nil {
		fmt.Printf("Unable to write to file: %s", *filePath)
	}
}

func toString(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
