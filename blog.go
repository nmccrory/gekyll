package main 

import(
	"log"
	"time"
	"sort"
	"io/ioutil"
	"encoding/json"
	"html/template"
)


// Type Safe creates objects for storing raw blog post data. 
type Safe struct {
	Route string
	Title string
	Author string
	Date string
	Body string
}

// Type Unsafe creates objects for storing blog post data 
// with a Body field holding an unescaped HTML template.
//
// This enables blog content with formatted HTML
// when rendered in Go template (e.g. code snippets, links).
type Unsafe struct {
	Route string
	Title string
	Author string
	Date time.Time
	Body template.HTML
}


/* ---Begin code for blog sorting--- */

// Type DateSorter is a data interface for the sort package to 
//  function calls to. 
type DateSorter []Safe

// Sort package makes three function calls to execute sort.Sort - 
// Len, Swap, and Less. Here these three functions are defined to 
// sort an array of DateSorter objects by their Date field. 
func (a DateSorter) Len() int           { return len(a) }
func (a DateSorter) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a DateSorter) Less(i, j int) bool { return a[i].Date > a[j].Date }


/* ---Functions for JSON I/O--- */

// The 'read' function reads a JSON file from a directory and unmarshals // the JSON into Go type Safe.
func read(f string) Safe {
	s := "./blogs/" + f + ".json"
	
	stream, err := ioutil.ReadFile(s)
	if err != nil {
		log.Println("Unsuccessful read while on file: ", f)
	}

	d := Safe{}
	json.Unmarshal(stream, &d)
	
	return d
}

// The 'readAll' function scans over every file in a directory and 
// unmarshals them from JSON objects to Go type Safe. Returns a slice of 
// type Safe with length equal to number of files discovered.
func readAll() []Safe {
	fs, err := ioutil.ReadDir("./blogs")
	if err != nil {
		log.Println("Could not locate directory to execute readAll.")
		log.Fatal(err)
	}

	sa := make([]Safe, len(fs))

	for i := range fs {
		s := "./blogs/" + fs[i].Name()
		f, _ := ioutil.ReadFile(s)

		r := Safe{}
		json.Unmarshal(f, &r)
		
		sa[i] = r
	}
	
	return sa
}

/* ---Data formatting functions--- */

// Converts type Safe to type Unsafe
func toUnsafe(a Safe) Unsafe {
	//converts Date to time.Time and Body to unsafe HTML
	u := Unsafe{Title: a.Title, Author: a.Author, Date: toTime(a.Date), Body: toHtml(a.Body)}
	return u
}

//Sorts the array of Safe objects by Date.
func sortByDate(safeArr []Safe) {
	sort.Sort(DateSorter(safeArr))
}

// Converts given string to HTML unsafe format
func toHtml(s string) template.HTML {
	return template.HTML(s)
}
// Converts Date string to Go type time.Time 
func toTime(t string) time.Time {
	p, err := time.Parse("01/02/2006", t)
	if err != nil {
		log.Println("Could not parse time: ", t)
	}

	return p
}