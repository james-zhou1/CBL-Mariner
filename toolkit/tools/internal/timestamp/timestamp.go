package timestamp

import (
	"log"
	"time"
)

// Call at the begining of the main() of a tool using "defer timestamp(time.Now(), "name_of_function")".
// The tool needs to import "time" too.
func Track(start time.Time, name string) {
	end := time.Now()
	diff := end.Sub(start)
	log.Printf("%s took %dms. Started at %s and ended at %s", name, diff.Nanoseconds()/1000, start, end)
}

// output as a string
// make a class output io.Writer
// features: initialize timestamp, flag each run (start & end, wipe out? ),
// go tool for csv files (for future parsing), tool name, step name, time, flag for time range
// output sth in the trace level?
// create struct that
// figure out logger package (how to call logger everywhere without passing a parameter)
