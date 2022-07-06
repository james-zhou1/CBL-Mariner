package csvparser

import "testing"

func Test_ParseCSV(t *testing.T) {
	files := FilepathsToArrayTest()
	ParseFiles(files)
}
