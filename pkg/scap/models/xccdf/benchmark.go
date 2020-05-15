package xccdf

import (
	"encoding/xml"
)

type Benchmark struct {
	XMLName  xml.Name `xml:Benchmark`
	Id       string   `xml:id,attr`
	Resolved string   `xml:resolved,attr`
	style    string   `xml:style,attr`
}
