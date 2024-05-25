package scrap

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestExtractText(t *testing.T) {

	bFile, err := os.ReadFile("resources/test_website.html")

	if err != nil {

		t.Errorf("could not read file because: %v", err)
		t.Fail()
	}

	newReader := bytes.NewReader(bFile)

	subset, err := DetectOcc(newReader, "test")

	if err != nil {

		t.Errorf("ExtractText failed because:( %v)", err)
		t.Fail()
	}

	bResult, err := io.ReadAll(subset)

	if err != nil {

		t.Errorf("could not read the rsult subset selected failed because:( %v)", err)
		t.Fail()
	}

	EXPECTED := `uction to test for p`

	if strings.Compare(EXPECTED, string(bResult)) != 0 {

		t.Errorf("not correct expected result failed becausse (%s) is not (%s), strings compare %v", string(bResult), EXPECTED, strings.Compare(EXPECTED, string(bResult)))
		t.Fail()
	}
}
