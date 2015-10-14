package filter

import (
	"strconv"
	"time"
)

//name - real name
//un - username
//alias -
func (f *Feed) Personal(name string, un string, alias string) {
	e := Entry{Title: "Mail Filter", Id: strconv.Itoa(f.Entrycount), Updated: time.Now(),
		Category: Category{Term: `filter`, Text: ""}}
	e.AddProperty("hasTheWord", "to:"+name+" OR to:"+un+" OR to:"+alias)
	e.AddProperty("shouldAlwaysMarkAsImportant", "true")
	e.AddProperty("shouldStar", "true")
	e.AddProperty("sizeOperator", "s_sl")
	e.AddProperty("sizeUnit", "s_smb")

	f.AddEntry(e)
}

func (f *Feed) Team() {
}

func (f *Feed) Environment() {
}

func (f *Feed) Entertaiment() {
}
