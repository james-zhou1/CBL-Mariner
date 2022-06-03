// Copyright Microsoft Corporation.
// Licensed under the MIT License.

package timestamp

import (
	"os"
	"testing"
	"time"
)

//TestMain found in configuration_test.go.

func Test_WritetoFile_Range_Instant(t *testing.T) {
	TrackToFile(time.Now(), "test tool", "test step", true, os.Stdout)
}

func Test_WritetoFile_noRange_Instant(t *testing.T) {
	TrackToFile(time.Now(), "test tool", "test step", false, os.Stdout)
}

func Test_WritetoCSV_Range_Instant(t *testing.T) {
	TrackToCSV(time.Now(), "test tool", "test step", true)
}
