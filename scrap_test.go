package scrap

import (
	"bytes"
	"html"
	"io"
	"os"
	"strings"
	"testing"
	"time"
)

func TestDetectOcc(t *testing.T) {

	start := time.Now()
	bFile, err := os.ReadFile("resources/test_website.html")

	if err != nil {
		t.Errorf("could not read file because: %v", err)
		t.Fail()
	}

	var newReader io.Reader = bytes.NewReader(bFile)

	subset, err := NewScrapper(10).DetectOcc(&newReader, "test")

	if err != nil {
		t.Errorf("ExtractText failed because:( %v)", err)
		t.Fail()
	}

	bResult, err := io.ReadAll(*subset)

	if err != nil {
		t.Errorf("could not read the rsult subset selected failed because:( %v)", err)
		t.Fail()
	}

	EXPECTED := `uction to test for p`

	if strings.Compare(EXPECTED, string(bResult)) != 0 {
		t.Errorf("not correct expected result failed becausse (%s) is not (%s), strings compare %v", string(bResult), EXPECTED, strings.Compare(EXPECTED, string(bResult)))
		t.Fail()
	}

	elapsed := time.Since(start)

	t.Log(elapsed)
}

func TestEscapedText(t *testing.T) {

	bFile, err := os.ReadFile("resources/test_website.html")

	if err != nil {

		t.Errorf("could not read file because: %v", err)
		t.Fail()
	}

	var escp string = html.EscapeString(string(bFile))

	var strReader io.Reader = strings.NewReader(escp)

	subset, err := NewScrapper(10).DetectOcc(&strReader, "test")

	if err != nil {
		t.Errorf("ExtractText failed because:( %v)", err)
		t.Fail()
	}

	bResult, err := io.ReadAll(*subset)

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

func TestStartIndex(t *testing.T) {

	bFile, err := os.ReadFile("resources/test_website.html")

	if err != nil {

		t.Errorf("could not read file because: %v", err)
		t.Fail()
	}

	var escp string = html.EscapeString(string(bFile))

	var strReader io.Reader = strings.NewReader(escp)

	subset, err := NewScrapper(20).DetectOcc(&strReader, "DOCTYPE")

	if err != nil {
		t.Errorf("ExtractText failed because:( %v)", err)
		t.Fail()
	}

	bResult, err := io.ReadAll(*subset)

	if err != nil {
		t.Errorf("could not read the rsult subset selected failed because:( %v)", err)
		t.Fail()
	}

	EXPECTED := `&lt;!DOCTYPE html&gt;&lt;`

	if strings.Compare(Clean(EXPECTED), Clean(string(bResult))) != 0 {
		t.Errorf("not correct expected result failed becausse (%s) is not (%s), strings compare %v", string(bResult), EXPECTED, strings.Compare(EXPECTED, string(bResult)))
		t.Fail()
	}
}

func TestEndIndex(t *testing.T) {

	bFile, err := os.ReadFile("resources/test_website.html")

	if err != nil {

		t.Errorf("could not read file because: %v", err)
		t.Fail()
	}

	var escp string = html.EscapeString(string(bFile))

	var strReader io.Reader = strings.NewReader(escp)

	subset, err := NewScrapper(5000).DetectOcc(&strReader, "https://dev.to/assets/signupM")

	if err != nil {
		t.Errorf("ExtractText failed because:( %v)", err)
		t.Fail()
	}

	bResult, err := io.ReadAll(*subset)

	if err != nil {
		t.Errorf("could not read the rsult subset selected failed because:( %v)", err)
		t.Fail()
	}

	if len(string(bResult)) > 5000*2 {
		t.Errorf("got more text than should have : %v", len(string(bResult)))
		t.Fail()
	}
}

func Clean(str string) string {

	return strings.ReplaceAll(str, "/n", "")

}
