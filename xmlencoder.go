package main

import (
	"encoding/xml"
	"fmt"
	"github.com/dasfoo/gmail-filters/filter"
	"github.com/dasfoo/gmail-filters/writer"
	"strconv"
	"time"
)

func main() {

	v := &filter.Feed{Title: "Mail Filters", Update: time.Now()}
	v.AddAuthor("", "")
	v.Xmlns = "http://www.w3.org/2005/Atom"
	v.Xmlnapp = "http://schemas.google.com/apps/2006"

	e := filter.Entry{Title: "Mail Filter", Id: strconv.Itoa(v.Entrycount), Updated: time.Now(),
		Category: filter.Category{Term: `filter`, Text: ""}}
	e.AddProperty("label", "nowork")
	v.AddEntry(e)

	v.Personal("", "", "")

	output, err := xml.MarshalIndent(v, "  ", "   ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println(xml.Header + string(output))
	writer.WriteToFile("test.xml", xml.Header+string(output))
}
