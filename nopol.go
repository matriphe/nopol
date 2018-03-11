package nopol

import (
	"fmt"
	"regexp"
	"strings"
)

// Pattern is regex pattern for Indonesian vehicle registration number
const Pattern = `^([A-Za-z]{1,2})(\s|-)*([0-9]{1,4})(\s|-)*([A-Za-z]{0,3})$`

// setPattern for regex of Indonesian police number
func setPattern() (r *regexp.Regexp, err error) {
	return regexp.Compile(Pattern)
}

// IsValid check if string is a valid Indonesian vehicle registration number
func IsValid(s string) bool {
	r, _ := setPattern()

	return r.MatchString(strings.TrimSpace(s))
}

// Format string to a valid Indonesian vehicle registration number
func Format(s string) (res string, err error) {
	s = strings.TrimSpace(s)

	if !IsValid(s) {
		return "", fmt.Errorf("%s is not a valid vehicle registration number format", s)
	}

	r, _ := setPattern()
	res = r.ReplaceAllString(s, `$1 $3 $5`)

	return strings.ToUpper(strings.TrimSpace(res)), nil
}
