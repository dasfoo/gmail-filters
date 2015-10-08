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

func (e *Entry) AddProperty(name string, value string) {
	pr := Property{Name: name, Value: value}
	e.Property = append(e.Property, pr)
}
