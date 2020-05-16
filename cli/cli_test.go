package cli

import "testing"

func TestCLI(t *testing.T) {
	
	result := HelloWorld()

	if result != "Hello. LibreSCAP is awesome." {
		t.Errorf("HelloWorld() = %s; want Hello. LibreSCAP is awesome.", result)
	}

}
