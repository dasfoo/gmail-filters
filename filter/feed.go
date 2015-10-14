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
	f.Entry = append(f.Entry, e)
	f.Entrycount++

}

type Feed struct {
	XMLName xml.Name  `xml:"feed"`
	Xmlns   string    `xml:"xmlns,attr"`
	Xmlnapp string    `xml:"xmlns:apps,attr"`
	Title   string    `xml:"title"`
	Update  time.Time `xml:"updated"`
	Author  Author
	Entry   []Entry

	Entrycount int `xml:"-"`
}
