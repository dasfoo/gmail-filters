package filter

import (
	"encoding/xml"
	"time"
)

type Author struct {
	XMLName xml.Name `xml:"author"`
	Name    string   `xml:"name"`
	Email   string   `xml:"email"`
}

type Category struct {
	Term string `xml:"term,attr"`
	Text string `xml:",chardata"`
}

func (f *Feed) AddAuthor(name string, email string) {
	f.Author = Author{Name: name, Email: email}
}

func (f *Feed) AddEntry(e Entry) {
	e.AddProperty("sizeOperator", "s_sl")
	e.AddProperty("sizeUnit", "s_smb")
	f.Entries = append(f.Entries, e)

}

type Feed struct {
	XMLName xml.Name  `xml:"feed"`
	Xmlns   string    `xml:"xmlns,attr"`
	Xmlnapp string    `xml:"xmlns:apps,attr"`
	Title   string    `xml:"title"`
	Update  time.Time `xml:"updated"`
	Author  Author
	Entries []Entry
}

func NewFeed() *Feed {
	f := Feed{}

	f.Xmlns = "http://www.w3.org/2005/Atom"
	f.Xmlnapp = "http://schemas.google.com/apps/2006"
	f.Update = time.Now()

	return &f

}
