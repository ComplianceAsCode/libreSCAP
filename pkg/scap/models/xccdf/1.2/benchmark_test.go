package xccdf_test

import (
	"testing"

	"github.com/complianceascode/librescap/pkg/scap/scap_document"
	"github.com/stretchr/testify/assert"
)

func TestRFCFeedParsing(t *testing.T) {
	doc, err := scap_document.ReadDocumentFromFile("../../../../../examples/scap/xccdf/1.2/test_xccdf_complex_check_nand.xccdf.xml")
	assert.Nil(t, err)
	assert.NotNil(t, doc.Benchmark)
	assert.Equal(t, doc.Benchmark.XMLName.Space, "http://checklists.nist.gov/xccdf/1.2")
}
