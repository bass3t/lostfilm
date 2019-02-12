package year

import (
	"strings"
)

const (
	// Year - "y"

	Y19702000 = "1970"
	Y20012005 = "2001"
	Y20062010 = "2006"
	Y20112015 = "2011"
	Y20162018 = "2016"
)

type FYear struct {
	value []string
}

func (f *FYear) Clear() {
	f.value = []string{}
}

func (f *FYear) Set(v string) {
	f.value = []string{v}
}

func (f *FYear) Add(v string) {
	f.value = append(f.value, v)
}

func (f *FYear) Value() string {
	if len(f.value) != 0 {
		return strings.Join(f.value, ",")
	}
	return ""
}
