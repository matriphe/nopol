package nopol

import (
	"errors"
	"regexp"
	"strings"
)

// R is regex pattern for Indonesian police number
const R = `^([A-Za-z]{1,2})(\s|-)*([0-9]{1,4})(\s|-)*([A-Za-z]{0,3})$`

// setPattern for regex of Indonesian police number
func setPattern() (r *regexp.Regexp, err error) {
	return regexp.Compile(R)
}

// IsValid check if string is a valid Indonesian police number
func IsValid(s string) bool {
	r, _ := setPattern()

	return r.MatchString(strings.TrimSpace(s))
}

// Format string to a valid Indonesian polica number
func Format(s string) (res string, err error) {
	s = strings.TrimSpace(s)

	if !IsValid(s) {
		return "", errors.New("Not a valid police number")
	}

	r, _ := setPattern()
	res = r.ReplaceAllString(s, `$1 $3 $5`)

	return strings.ToUpper(strings.TrimSpace(res)), nil
}
