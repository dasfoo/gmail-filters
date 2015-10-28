package main

import (
	"encoding/xml"
	"fmt"
	"github.com/dasfoo/gmail-filters/filter"
)

func main() {

	v := filter.NewFeed()
	v.AddAuthor("Ivan Ivanov", "ivan@mail.ru")

	v.Personal([]string{"Ivan", "Ivanov"},
		[]string{"ivan.ivanov", "ivanov", "iv"})
	v.Team("team1@example.org", "team2@example.org", "team3@example.org")

	v.Environment("go.example.com", "kafka.example.com")
	v.Entertainment("english.example.com", "shopping.example.com")
	v.Project("GmailFilters", "it is interesting", true)
	v.Project("DasFoo", "interesting", false)

	output, err := xml.MarshalIndent(v, "  ", "   ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println(xml.Header + string(output))
}
