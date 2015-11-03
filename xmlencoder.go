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
	v.Review("Projectname", "xtrareview.example.com", "plusrev.example.com")
	v.Project("Gmaillabel", "GmailFilters", "it is interesting", true)
	v.Project("projectlabel", "DasFoo", "interesting", false)
	v.Service("Service", "service1.example.com", "service2.example.com")
	v.AutomatedSystem("auto", "auto1@example.com", "auto2@example.com")

	output, err := xml.MarshalIndent(v, "  ", "   ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Println(xml.Header + string(output))
}
