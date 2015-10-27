package filter

import (
	"log"
	"strconv"
	"strings"
	"time"
)

func (f *Feed) Personal(real_name []string, emails []string) {

	sname := strings.Join(real_name, " OR ")
	smails := strings.Join(emails, " OR ")
	e := Entry{Title: "Mail Filter", Id: "tag:dasfoo.filters,smartfilter:personal", Updated: time.Now(),
		Category: Category{Term: `filter`, Text: ""}}
	e.AddProperty("hasTheWord", sname+" OR "+smails)
	e.AddProperty("to", smails)
	e.AddProperty("shouldAlwaysMarkAsImportant", "true")
	e.AddProperty("shouldStar", "true")

	f.AddEntry(e)
}

func (f *Feed) Team(mails ...string) {
	s := strings.Join(mails, " OR ")
	log.Println(s)
	e := Entry{Title: "Mail Filter", Id: "tag:dasfoo.filters,smartfilter:team", Updated: time.Now(),
		Category: Category{Term: `filter`, Text: ""}}
	e.AddProperty("from", s)
	e.AddProperty("label", "team")
	e.AddProperty("shouldAlwaysMarkAsImportant", "true")

	f.AddEntry(e)

}

func (f *Feed) Environment(listnames ...string) {
	s := strings.Join(listnames, "> OR <")
	s = "list:(<" + s + ">)"
	e := Entry{Title: "Mail Filter", Id: "tag:dasfoo.filters,smartfilter:environment", Updated: time.Now(),
		Category: Category{Term: `filter`, Text: ""}}
	e.AddProperty("hasTheWord", s)
	e.AddProperty("label", "Environment")
	e.AddProperty("shouldArchive", "true")

	f.AddEntry(e)
}

func (f *Feed) Entertainment(listnames ...string) {
	s := strings.Join(listnames, "> OR <")
	s = "list:(<" + s + ">)"
	e := Entry{Title: "Mail Filter", Id: "tag:dasfoo.filters,smartfilter:entertainment",
		Updated: time.Now(), Category: Category{Term: `filter`, Text: ""}}
	e.AddProperty("hasTheWord", s)
	e.AddProperty("label", "Entertainment")
	e.AddProperty("shouldArchive", "true")

	f.AddEntry(e)
}

func (f *Feed) TeamMessages(proj_not_intr bool, listnames ...string) {
	s := strings.Join(listnames, "> OR <")
	s = "list:(<" + s + ">)"
	e := Entry{Title: "Mail Filter",
		Id:      "tag:dasfoo.filters,smartfilter:teammes" + strconv.Itoa(len(f.Entries)),
		Updated: time.Now(), Category: Category{Term: `filter`, Text: ""}}
	e.AddProperty("hasTheWord", s)
	e.AddProperty("label", "Project")
	if proj_not_intr {
		e.AddProperty("shouldArchive", "true")
	}

	f.AddEntry(e)
}
