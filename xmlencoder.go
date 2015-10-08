package main

import (
	"encoding/xml"
	"fmt"
	"github.com/dasfoo/gmail-filters/filter"
	"os"
	"time"
)

func main() {

	v := &filter.Feed{Title: "Mail Filters", Update: time.Now()}
	v.AddAuthor("Artem Sheremet", "dotdoom@gmail.com")

	e := filter.Entry{Title: "Mail Filter", Id: "1234", Updated: time.Now(),
		Category: filter.Category{Term: `filter`, Text: ""}}
	e.AddProperty("label", "nowork")
	v.AddEntry(e)

	output, err := xml.MarshalIndent(v, "  ", "   ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write(output)
}
