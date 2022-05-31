// Copyright Microsoft Corporation.
// Licensed under the MIT License.

package timestamp

import (
	"testing"
	"time"
)

//TestMain found in configuration_test.go.

func Test_test(t *testing.T) {
	Track(time.Now(), "test file")
}
