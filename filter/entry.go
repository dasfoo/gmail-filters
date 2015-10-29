package filter

import (
	"encoding/xml"
	"time"
)

type Entry struct {
	XMLName  xml.Name   `xml:"entry"`
	Category Category   `xml:"category"`
	Title    string     `xml:"title"`
	Id       string     `xml:"id"`
	Updated  time.Time  `xml:"updated"`
	Content  string     `xml:"content"`
	Property []Property `xml:"apps:property"`
}

func NewEntry(title, id string) *Entry {
	entry := Entry{
		Title:    title,
		Id:       id,
		Updated:  time.Now(),
		Category: Category{Term: `filter`, Text: ""},
	}
	return &entry
}

func (e *Entry) AddProperty(name string, value string) {
	pr := Property{Name: name, Value: value}
	e.Property = append(e.Property, pr)
}
