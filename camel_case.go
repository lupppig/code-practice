package main

import (
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
