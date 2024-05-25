package scrap

import (
	"fmt"
	"html"
	"io"
	"os"
	"strconv"
	"strings"
)

type ScrapperHTML interface {
	ExtractText()
}

type Subset struct {
	Selected io.Reader
}

func ExtractText(r io.Reader, search string) (*Subset, error) {

	b, err := io.ReadAll(r)

	if err != nil {

		return nil, fmt.Errorf("cant read input from reader")
	}

	str := html.EscapeString(string(b))

	routines := os.Getenv("max_procs")
	i, err := strconv.Atoi(routines)

	if err != nil {
		return nil, fmt.Errorf("no mac_procs env variable set to correct format")
	}
	chars := []rune(str)

	routineTextSize := len(chars) / i

	startIndex := 0
	endIndex := routineTextSize

	//amount of text parts
	for index := range i {

		reader := strings.NewReader(string(chars[startIndex:endIndex]))

		//how to turnover result? need channel?
		detectOcc(reader, search)

		startIndex = routineTextSize

		endIndex = index * routineTextSize

	}

	return &Subset{}, nil

	//w8 for all routines to finish (the cancel each other out)

}

func detectOcc(search io.Reader, searchWord string) (io.Reader, error) {

	buf := new(strings.Builder)
	_, err := io.Copy(buf, search)

	if err != nil {

		return nil, fmt.Errorf("%s", "cant read from io Reader")
	}

	//place in the text where word was found
	foundIndex := strings.Index(buf.String(), searchWord)
	beforeIn := foundIndex - 200
	aftertIndex := foundIndex + 200

	surround := string([]rune(buf.String())[beforeIn:aftertIndex])

	reader := strings.NewReader(surround)

	if foundIndex != 0 {

		return reader, nil

	}

	return reader, nil

}
