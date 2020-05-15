package scap_document

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"

	"github.com/complianceascode/librescap/pkg/scap/models/xccdf"
)

const (
	xccdfBenchmarkElement = "Benchmark"
)

type Document struct {
	XMLName xml.Name `json:"-"`
	*xccdf.Benchmark
}

func ReadDocument(r io.Reader) (*Document, error) {
	d := xml.NewDecoder(r)
	for {
		token, err := d.Token()
		if err != nil || token == nil {
			return nil, fmt.Errorf("Could not decode XML: %v", err)
		}
		switch startElement := token.(type) {
		case xml.StartElement:
			switch startElement.Name.Local {
			case xccdfBenchmarkElement:
				var bench xccdf.Benchmark
				if err := d.DecodeElement(&bench, &startElement); err != nil {
					return nil, err
				}
				return &Document{Benchmark: &bench}, nil
			}
		}
	}
	return nil, fmt.Errorf("Could not parse input file")
}

func ReadDocumentFromFile(filepath string) (*Document, error) {
	reader, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	return ReadDocument(reader)
}
