package main

import (
	"regexp"
	"strings"
)

func ToCamelCase(s string) string {
	var str string
	var upNext bool

	for _, st := range s {
		ss := string(st)

		if ss == "-" || ss == "_" {
			upNext = true
			continue
		}

		if upNext {
			str += strings.ToUpper(ss)
			upNext = false
			continue
		}
		str += ss
	}

	return str
}

func ToCamelCase2(s string) string {
	reg := regexp.MustCompile("[-_](.)")
	return reg.ReplaceAllStringFunc(s, func(w string) string {
		return strings.ToUpper(string(w[1:]))
	})
}
