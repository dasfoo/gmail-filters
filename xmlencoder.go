package main

import (
	"encoding/xml"
	"fmt"
	"github.com/dasfoo/gmail-filters/filter"
	"time"
)

func main() {

	v := &filter.Feed{Title: "Mail Filters", Update: time.Now()}
	v.AddAuthor("Ivan Ivanov", "ivan@mail.ru")
	v.Xmlns = "http://www.w3.org/2005/Atom"
	v.Xmlnapp = "http://schemas.google.com/apps/2006"

	v.Personal([]string{"Ivan", "Ivanov"},
		[]string{"ivan.ivanov", "ivanov", "ivan"})
	v.Team("team1@example.org", "team2@example.org", "team3@example.org")

	v.Environment("g.example.com")
	v.Entertainment("aa.example.com")

	output, err := xml.MarshalIndent(v, "  ", "   ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println(xml.Header + string(output))
}
