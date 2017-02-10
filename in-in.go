package main

import (
	"github.com/atotto/clipboard"
	"regexp"
	"strings"
)

func main() {
	var err error
	var inClauseLimit = 999

	out, err := clipboard.ReadAll()
	if err != nil {
		panic(err)
	}

	var splitted = regexp.MustCompile("\r\n|\n\r|\n|\r").Split(out, -1)
	var mapped = Map(splitted, func(v string) string {
		return strings.TrimSpace(v)
	})
	var filtered = Filter(mapped, func(v string) bool {
		return len(strings.TrimSpace(v)) > 0
	})

	results := []string{}
	for in := range Chunks(filtered, inClauseLimit) {
		results = append(results, toStringInClause(in))
	}

	result := strings.Join(results, " or ")

	if err := clipboard.WriteAll(result); err != nil {
		panic(err)
	}
}

func toStringInClause(splitted []string) string {
	var result = " in ("
	for i, v := range splitted {
		var trimmed = strings.TrimSpace(v)
		if len(trimmed) > 0 {
			result = result + "'" + strings.TrimSpace(v) + "'"
		}

		if isEndOfInClause(splitted, i) == false {
			result = result + ", "
		}

	}
	return result + ")\r\n"
}

func isEndOfInClause(inClauseElement []string, index int) bool {
	return (len(inClauseElement) - 2 == index)
}

func Chunks(l []string, n int) chan []string {
	ch := make(chan []string)

	go func() {
		for i := 0; i < len(l); i += n {
			from_idx := i
			to_idx := i + n
			if to_idx > len(l) {
				to_idx = len(l)
			}
			ch <- l[from_idx:to_idx]
		}
		close(ch)
	}()
	return ch
}

func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}