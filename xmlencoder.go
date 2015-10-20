package main

import (
	"encoding/xml"
	"fmt"
	"github.com/dasfoo/gmail-filters/filter"
	"github.com/dasfoo/gmail-filters/writer"
	"time"
)

func main() {

	v := &filter.Feed{Title: "Mail Filters", Update: time.Now()}
	v.AddAuthor("Ivan Ivanov", "ivan@mail.ru")
	v.Xmlns = "http://www.w3.org/2005/Atom"
	v.Xmlnapp = "http://schemas.google.com/apps/2006"

	e := filter.Entry{Title: "Mail Filter", Id: strconv.Itoa(v.Entrycount), Updated: time.Now(),
		Category: filter.Category{Term: `filter`, Text: ""}}
	e.AddProperty("label", "nowork")
	v.AddEntry(e)

	v.Personal([]string{"Ivan", "Ivanov"},
		[]string{"ivan.ivanov", "ivanov", "ivan"},
		[]string{"iiv", "iv"})
	v.Team("team1@example.org", "team2@example.org", "team3@example.org")

	v.Environment("g.example.com")
	v.Entertainment("aa.example.com")

	output, err := xml.MarshalIndent(v, "  ", "   ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println(xml.Header + string(output))
	writer.WriteToFile("test.xml", xml.Header+string(output))
}
