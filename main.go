package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/reiver/go-path"
)

var location string
var verbose bool
var when string
var trial bool

func init() {
	flag.StringVar(&location, "location", "", "location; the year, month, and time can be different depending on the location. ex: --location=America/Vancouver or --location=Asia/Seoul or --location=Asia/Tehran")
	flag.BoolVar(&verbose, "v", false, "verbose")
	flag.StringVar(&when, "when", "", "when. ex: --when=2022-11-04T18:03:45Z")
	flag.BoolVar(&verbose, "t", false, "trial")

	flag.Parse()
}

// datetimelayout is the output date-time layout.
// This is in not parsing date-time layout.
const datetimelayout string = "2006-01-02T15:04:05-07:00"

func main() {

	var loc *time.Location
	{
		var err error

		loc, err = time.LoadLocation(location)
		if nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: could not load location %q: %s\n", location, err)
			os.Exit(1)
			return
		}
	}
	if verbose {
		fmt.Printf("--location=%s\n", loc)
		fmt.Printf("Location:  %s\n", loc)
	}

	var t time.Time
	{
		t = time.Now()

		if "" != when {
			var err error

			t, err = time.Parse(time.RFC3339, when)
			if nil != err {
				fmt.Fprintf(os.Stderr, "ERROR: could not parse date-time %q: %s\n", when, err)
				os.Exit(1)
				return
			}
		}

		t = t.In(loc)

	}
	if verbose {
		fmt.Printf("Date-Time: %s\n", t.Format(datetimelayout))
	}

	var dirname string
	{
		dirname += "log/"

		{
			var year int = t.Year()
			dirname += fmt.Sprintf("%d", year)
		}

		dirname += "/"

		{
			var month int = int(t.Month())
			if month < 10 {
				dirname += "0"
			}
			dirname += fmt.Sprintf("%d", month)
		}

		dirname += "/"

		{
			var day int = t.Day()
			if day < 10 {
				dirname += "0"
			}
			dirname += fmt.Sprintf("%d", day)
		}

	}
	if verbose {
		fmt.Printf("Dir-Name: %s\n", dirname)
	}

	var unixtimestamp int64 = t.Unix()
	if verbose {
		fmt.Printf("Unix-Time-Stamp: %d\n", unixtimestamp)
	}

	var filename string = fmt.Sprintf("%d.wiki", unixtimestamp)
	if verbose {
		fmt.Printf("File-Name: %s\n", filename)
	}

	var wikipath string = path.Join(dirname, filename)
	if verbose {
		fmt.Printf("Path: %s\n", wikipath)
	}

	if !trial {
		const permissions os.FileMode = 0755

		var path string = dirname

		err := os.MkdirAll(path, permissions)
		if nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: could not create directory %q: %s\n", path, err)
			os.Exit(1)
			return
		}

		fmt.Printf("created directory: %s\n", path)
	}

	if !trial {
		const permissions os.FileMode = 0644

		var path string = wikipath

		var data []byte = []byte(
			"wiki/1"+"\n"+
			""+"\n"+
			"§ Hello World!"+"\n"+
			""+"\n"+
			"⸺ by Joe Blow"+"\n"+
			""+"\n"+
			"⸺ published "+t.Format(datetimelayout)+"\n"+
			""+"\n"+
			"Hello world!"+"\n",
		)

		err := os.WriteFile(path, data, permissions)
		if nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: could not create file %q: %s\n", path, err)
			os.Exit(1)
			return
		}

		fmt.Printf("created file:      %s\n", path)
	}
}