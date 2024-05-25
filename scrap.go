package scrap

import (
	"fmt"
	"io"
	"strings"
)

type ScrapperHTML interface {
	ExtractText()
}

type Subset struct {
	Selected io.Reader
}

type Config struct {
	VisionSize int
}

type Scrapper struct {
	cfg Config
}

// func ExtractText(r io.Reader, search string) (*Subset, error) {

// 	b, err := io.ReadAll(r)

// 	if err != nil {

// 		return nil, fmt.Errorf("cant read input from reader")
// 	}

// 	str := html.EscapeString(string(b))

// 	if err != nil {
// 		return nil, fmt.Errorf("no mac_procs env variable set to correct format")
// 	}
// 	chars := []rune(str)

// 	routineTextSize := len(chars) / routines

// 	startIndex := 0
// 	endIndex := routineTextSize

// 	//amount of text parts
// 	for index := range routines {

// 		reader := strings.NewReader(string(chars[startIndex:endIndex]))

// 		//how to turnover result? need channel?
// 		resultReader, err := DetectOcc(reader, search)

// 		if err != nil {

// 			return nil, fmt.Errorf("detect Occurence function failed because %v")
// 		}

// 		startIndex = routineTextSize

// 		endIndex = index * routineTextSize

// 	}

// 	return &Subset{}, nil

// 	//w8 for all routines to finish (the cancel each other out)

// }

func NewScrapper(visionSize int) *Scrapper {

	return &Scrapper{cfg: Config{VisionSize: visionSize}}
}

func (s *Scrapper) DetectOcc(search *io.Reader, searchWord string) (*io.Reader, error) {

	if search == nil {
		return nil, fmt.Errorf("the io Reader is nil")
	}

	bRes, err := io.ReadAll(*search)
	if err != nil {
		return nil, fmt.Errorf("%s", "cant read from io Reader")
	}

	//place in the text where word was found
	foundIndex := strings.Index(string(bRes), searchWord)
	beforeIn := max(foundIndex-s.cfg.VisionSize, 0)
	aftertIndex := min(foundIndex+s.cfg.VisionSize, len(string(bRes)))

	surround := string([]rune(string(bRes))[beforeIn:aftertIndex])

	var reader io.Reader = strings.NewReader(surround)

	if foundIndex != 0 {
		return &reader, nil
	}

	return nil, nil

}
