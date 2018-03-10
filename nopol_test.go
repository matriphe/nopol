package nopol

import (
	"testing"
)

// Nopol struct
type Nopol struct {
	s string
	b bool
	r string
}

var nopols = []Nopol{
	{"AD 6742 DZ", true, "AD 6742 DZ"},
	{"AD6742DZ", true, "AD 6742 DZ"},
	{"AD-6742-DZ", true, "AD 6742 DZ"},
	{"     AD-6742 DZ		", true, "AD 6742 DZ"},
	{"B 1 XYZ", true, "B 1 XYZ"},
	{"RI 1", true, "RI 1"},
	{"CD 123", true, "CD 123"},
	{"12345", false, ""},
	{"AB 1234567890 CD", false, ""},
	{"AB 1 WXYZ", false, ""},
	{"AB CDEF GHI", false, ""},
}

func TestIsValid(t *testing.T) {
	for _, n := range nopols {
		r := IsValid(n.s)

		if n.b != r {
			t.Errorf("Police number %s is not valid, got %t, want %t", n.s, r, n.b)
		}
	}
}

func TestFormat(t *testing.T) {
	for _, n := range nopols {
		r, _ := Format(n.s)

		if n.r != r {
			t.Errorf("Police number %s is not formatted, got %s, want %s", n.s, r, n.r)
		}
	}
}
