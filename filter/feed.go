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

}

type Feed struct {
	XMLName xml.Name `xml:"feed"`

	Title  string    `xml:"title"`
	Update time.Time `xml:"update"`
	Author Author
	Entry  []Entry
}
