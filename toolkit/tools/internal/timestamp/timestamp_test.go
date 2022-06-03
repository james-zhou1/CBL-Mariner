// Copyright Microsoft Corporation.
// Licensed under the MIT License.

package timestamp

import (
	"os"
	"testing"
	"time"
)

//TestMain found in configuration_test.go.

func Test_WritetoFile_range_instant(t *testing.T) {
	TrackToFile(time.Now(), "test tool", "test step", true, os.Stdout)
}

func Test_WritetoFile_noRange_instant(t *testing.T) {
	TrackToFile(time.Now(), "test tool", "test step", false, os.Stdout)
}

func Test_WritetoCSV_range_instant(t *testing.T) {
	TrackToCSV(time.Now(), "test tool", "test step", true)
}

func Test_WritetoCSV_noRange_instant(t *testing.T) {
	TrackToCSV(time.Now(), "test tool", "test step", false)
}

func Test_WritetoFile_range_sleeps(t *testing.T) {
	defer TrackToFile(time.Now(), "tool 1", "step 1", true, os.Stdout)
	time.Sleep(3 * time.Second)
}

func Test_WritetoCSV_range_sleeps(t *testing.T) {
	defer TrackToCSV(time.Now(), "test tool", "test step", true)
	time.Sleep(3 * time.Second)
}
