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

func (f *Feed) AddEntry(entry *Entry) {
	entry.AddProperty("sizeOperator", "s_sl")
	entry.AddProperty("sizeUnit", "s_smb")
	f.Entries = append(f.Entries, *entry)

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
	feed := Feed{}

	feed.Xmlns = "http://www.w3.org/2005/Atom"
	feed.Xmlnapp = "http://schemas.google.com/apps/2006"
	feed.Update = time.Now()

	return &feed

}
