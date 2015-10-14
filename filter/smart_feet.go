package filter

import (
	"log"
	"strconv"
	"strings"
	"time"
)

//name - real name
//un - username
//alias -
func (f *Feed) Personal(name []string, un []string, alias []string) {
	ressl := append(un, alias...)
	s := strings.Join(ressl, " OR to:")
	sname := strings.Join(name, " OR ")
	e := Entry{Title: "Mail Filter", Id: strconv.Itoa(f.Entrycount), Updated: time.Now(),
		Category: Category{Term: `filter`, Text: ""}}
	e.AddProperty("hasTheWord", sname+" OR to:"+s)
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

func (f *Feed) Environment() {
}

func (f *Feed) Entertaiment() {
}
