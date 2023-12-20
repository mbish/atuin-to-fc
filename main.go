package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"
)

func dropNull(data []byte) []byte {
	if len(data) > 0 && data[len(data)-1] == '\000' {
		return data[0 : len(data)-1]
	}
	return data
}

func ScanNulls(data []byte, atEOF bool) (advance int, token []byte, err error) {
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}
	if i := bytes.IndexByte(data, '\000'); i >= 0 {
		return i + 1, dropNull(data[0:i]), nil
	}
	if atEOF {
		return len(data), dropNull(data), nil
	}
	// Request more data.
	return 0, nil, nil
}

func printResult(maxLength int, index int, text string) {
	paddingLength := maxLength - len(fmt.Sprintf("%d", index))
	padding := strings.Repeat(" ", paddingLength)
	fmt.Printf(padding+"%d  %s\n", index, text)
}

func main() {
	var reverse bool

	flag.BoolVar(&reverse, "r", false, "Print entries in reverse order")
	flag.Parse()

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(ScanNulls)
	var results []string

	for scanner.Scan() {
		rawText := scanner.Text()
		text := strings.Replace(rawText, "\n", "\\n", -1)
		results = append(results, text)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	maxLength := len(fmt.Sprintf("%d", len(results)))

	if reverse {
		for i := len(results) - 1; i > 0; i-- {
			printResult(maxLength, i, results[i])
		}
	} else {
		for i := 0; i < len(results); i++ {
			printResult(maxLength, i, results[i])
		}
	}

}
