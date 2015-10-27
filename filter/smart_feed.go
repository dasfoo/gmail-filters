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

func (f *Feed) Project(projectname, intr_str string, proj_not_intr bool) {
	s := projectname + "-users> OR <" + projectname + "-team> OR <" + projectname + "-sre"
	s = "list:(<" + s + ">)"

	mess := Entry{
		Title:    "Mail Filter",
		Id:       "tag:dasfoo.filters,smartfilter:project" + strconv.Itoa(len(f.Entries)),
		Updated:  time.Now(),
		Category: Category{Term: `filter`, Text: ""}}
	mess.AddProperty("hasTheWord", s)
	mess.AddProperty("label", projectname)
	if proj_not_intr {
		mess.AddProperty("shouldArchive", "true")
	}
	f.AddEntry(mess)

	s = projectname + "-reviews OR " + projectname + "+reviews"

	rev := Entry{
		Title:    "Mail Filter",
		Id:       "tag:dasfoo.filters,smartfilter:project" + strconv.Itoa(len(f.Entries)),
		Updated:  time.Now(),
		Category: Category{Term: `filter`, Text: ""}}
	rev.AddProperty("to", s)
	rev.AddProperty("shouldArchive", "true")
	f.AddEntry(rev)

	intr := Entry{
		Title:    "Mail Filter",
		Id:       "tag:dasfoo.filters,smartfilter:project" + strconv.Itoa(len(f.Entries)),
		Updated:  time.Now(),
		Category: Category{Term: `filter`, Text: ""}}
	intr.AddProperty("to", s)
	intr.AddProperty("hasTheWord", intr_str)
	intr.AddProperty("label", projectname)
	intr.AddProperty("shouldAlwaysMarkAsImportant", "true")
	f.AddEntry(intr)

}
