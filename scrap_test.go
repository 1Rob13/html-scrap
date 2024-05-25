package scrap

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestExtractText(t *testing.T) {

	bFile, err := os.ReadFile("resources/test_website.html")

	if err != nil {

		t.Errorf("could not read file because: %v", err)
		t.Fail()
	}

	newReader := bytes.NewReader(bFile)

	subset, err := ExtractText(newReader, "test")

	if err != nil {

		t.Errorf("ExtractText failed because:( %v)", err)
		t.Fail()
	}

	bResult, err := io.ReadAll(subset.Selected)

	if err != nil {

		t.Errorf("could not read the rsult subset selected failed because:( %v)", err)
		t.Fail()
	}

	EXPECTED := "hi this is a test"

	if string(bResult) != "hi this is a test" {

		t.Errorf("not correct expected result failed becausse (%s) is not (%s)", string(bResult), EXPECTED)
		t.Fail()
	}
}
