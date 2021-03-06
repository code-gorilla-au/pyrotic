package formats

import (
	"regexp"
	"strings"
	"unicode"
)

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap = regexp.MustCompile("([a-z0-9])([A-Z])")
var matchSymbol = regexp.MustCompile("[_-]")

func CaseSnake(str string) string {
	tmp := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	tmp = matchAllCap.ReplaceAllString(tmp, "${1}_${2}")
	tmp = matchSymbol.ReplaceAllString(tmp, "${1}_${2}")
	return strings.ToLower(tmp)
}

func CasePascal(str string) string {
	tmp := matchFirstCap.ReplaceAllString(str, "${1} ${2}")
	tmp = matchSymbol.ReplaceAllString(tmp, "${1} ${2}")
	tmp = strings.Title(tmp)
	return strings.ReplaceAll(tmp, " ", "")
}

func CaseKebab(str string) string {
	tmp := matchFirstCap.ReplaceAllString(str, "${1}-${2}")
	tmp = matchAllCap.ReplaceAllString(tmp, "${1}-${2}")
	tmp = matchSymbol.ReplaceAllString(tmp, "${1}-${2}")
	return strings.ToLower(tmp)
}

func CaseCamel(str string) string {
	tmp := matchFirstCap.ReplaceAllString(str, "${1} ${2}")
	tmp = matchSymbol.ReplaceAllString(tmp, "${1} ${2}")
	tmp = strings.Title(tmp)
	tmp = strings.ReplaceAll(tmp, " ", "")
	return lowercaseFirst(tmp)
}

func lowercaseFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
