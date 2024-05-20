package scrap

import (
	"context"
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

type Scrapper struct {
}

type Text struct {
	Occurences uint8
	Selected   []string
	All        io.Reader
}

func (s *Scrapper) ExtractText(r io.Reader, search string) (*Text, error) {

	b, err := io.ReadAll(r)

	if err != nil {

		return nil, fmt.Errorf("cant read input from reader")
	}

	str := html.EscapeString(string(b))

	if err != nil {
		return nil, fmt.Errorf("no mac_procs env variable set to correct format")
	}

	routines := os.Getenv("max_procs")
	i, err := strconv.Atoi(routines)

	chars := []rune(str)

	routineTextSize := len(chars) / i

	startIndex := 0
	endIndex := routineTextSize

	ctxB := context.Background()
	ctx, cancel := context.WithCancel(ctxB)
	//amount of text parts
	for index := range i {

		//how to turnover result? need channel?
		go detectOcc(ctx, cancel, string(chars[startIndex:endIndex]), search)

		startIndex = routineTextSize

		endIndex = index * routineTextSize

	}

	//w8 for all routines to finish (the cancel each other out)

}

func detectOcc(context context.Context, cancel context.CancelFunc, search string, searchWord string) Text {

	//place in the text where word was found
	strings.Index(search, searchWord)

	//get last line break

	//get next line break

	//copy all to io.Reader

	//stop detect if already found somewhere
	select {}

}
