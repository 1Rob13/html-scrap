package scrap

import (
	"bytes"
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
	Selected io.Reader
	All      io.Reader
}

func (s *Scrapper) ExtractText(r io.Reader, search string) (*Text, error) {

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

	ctxB := context.Background()
	ctx, cancel := context.WithCancel(ctxB)
	sigs := make(chan Text, 1)
	//amount of text parts
	for index := range i {

		//how to turnover result? need channel?
		go detectOcc(ctx, cancel, sigs, string(chars[startIndex:endIndex]), search)

		startIndex = routineTextSize

		endIndex = index * routineTextSize

	}

	msg := <-sigs
	return &msg, nil

	//w8 for all routines to finish (the cancel each other out)

}

func detectOcc(context context.Context, cancelCall context.CancelFunc, sig chan Text, search string, searchWord string) {

	//place in the text where word was found
	foundIndex := strings.Index(search, searchWord)

	//get last line break

	beforeIn := foundIndex - 200
	aftertIndex := foundIndex + 200

	surround = string([]rune(search)[beforeIn:aftertIndex])

	io.Reader.Read([]byte(surround))

	//get next line break

	//copy all to io.Reader

	if foundIndex != 0 {

		sig <- Text{ Selected: }

		cancelCall()
	}

}
