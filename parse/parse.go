package parse

import (
	"fmt"
	"strconv"
	"strings"
)

type ParseError struct {
	index int
	word  string
	err   error
}

func (e *ParseError) String() string {
	return fmt.Sprintf("index: %d ,word: %s .error %v", e.index, e.word, e.err)
}

func Parse(line string) (numbers []int, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("in parse func:  %v", r)
			}
		}
	}()
	fields := strings.Fields(line)
	numbers = stringsToNumbers(fields)
	return
}

func stringsToNumbers(fields []string) (intArr []int) {
	if len(fields) == 0 {
		panic("no words to parse")
	}
	intArr = make([]int, 0)
	for index, ele := range fields {
		int, err := strconv.Atoi(ele)
		if err != nil {
			panic(&ParseError{index, ele, err})
		}
		intArr = append(intArr, int)
	}
	return
}
