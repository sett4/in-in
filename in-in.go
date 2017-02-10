package main

import (
	"github.com/atotto/clipboard"
	"regexp"
	"strings"
)

func main() {
	var err error

	out, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}

	var result = ""
	var splitted = regexp.MustCompile("\r\n|\n\r|\n|\r").Split(out, -1)
	for i, v := range splitted {
		var trimmed = strings.TrimSpace(v)
		if len(trimmed) > 0 {
			result = result + "'" + strings.TrimSpace(v) + "'"
		}

		if len(splitted) - 2 > i {
			result = result + ",\n"
		}

	}

	if err := clipboard.WriteAll(result); err != nil {
		panic(err)
	}
}
