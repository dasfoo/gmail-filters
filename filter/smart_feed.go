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
	e := Entry{Title: "Mail Filter", Id: strconv.Itoa(f.Entrycount), Updated: time.Now(),
		Category: Category{Term: `filter`, Text: ""}}
	e.AddProperty("hasTheWord", sname)
	e.AddProperty("to", smails)
	e.AddProperty("shouldAlwaysMarkAsImportant", "true")
	e.AddProperty("shouldStar", "true")
	e.AddProperty("sizeOperator", "s_sl")
	e.AddProperty("sizeUnit", "s_smb")

	f.AddEntry(e)
}

func (f *Feed) Team(mails ...string) {
	s := strings.Join(mails, " OR ")
	log.Println(s)
	e := Entry{Title: "Mail Filter", Id: strconv.Itoa(f.Entrycount), Updated: time.Now(),
		Category: Category{Term: `filter`, Text: ""}}
	e.AddProperty("from", s)
	e.AddProperty("label", "team")
	e.AddProperty("shouldAlwaysMarkAsImportant", "true")
	e.AddProperty("sizeOperator", "s_sl")
	e.AddProperty("sizeUnit", "s_smb")

	f.AddEntry(e)

}

func (f *Feed) Environment(listnames ...string) {
	s := strings.Join(listnames, ">) OR (<")
	s = "list:(<" + s + ">)"
	e := Entry{Title: "Mail Filter", Id: strconv.Itoa(f.Entrycount), Updated: time.Now(),
		Category: Category{Term: `filter`, Text: ""}}
	e.AddProperty("hasTheWord", s)
	e.AddProperty("label", "Environment")
	e.AddProperty("shouldArchive", "true")
	e.AddProperty("sizeOperator", "s_sl")
	e.AddProperty("sizeUnit", "s_smb")

	f.AddEntry(e)
}

func (f *Feed) Entertainment(listnames ...string) {
	s := strings.Join(listnames, ">) OR (<")
	s = "list:(<" + s + ">)"
	e := Entry{Title: "Mail Filter", Id: strconv.Itoa(f.Entrycount), Updated: time.Now(),
		Category: Category{Term: `filter`, Text: ""}}
	e.AddProperty("hasTheWord", s)
	e.AddProperty("label", "Entertainment")
	e.AddProperty("shouldArchive", "true")
	e.AddProperty("sizeOperator", "s_sl")
	e.AddProperty("sizeUnit", "s_smb")

	f.AddEntry(e)
}
