package genre

import (
	"strings"
)

const (
	// Genre - "g"

	Comedy    = "1"
	Drama     = "2"
	Fiction   = "3"
	Thriller  = "4"
	Detective = "5"
	Mysticism = "6"
	Action    = "7"
	Western   = "8"
	Adventure = "9"
	Horror    = "11"
	Fantasy   = "12"
	Crime     = "13"
	History   = "14"
	Biography = "17"
	War       = "18"
	Comics    = "19"
	Family    = "20"
	Melodra   = "21"
)

type FGenre struct {
	value []string
}

func (f *FGenre) Clear() {
	f.value = []string{}
}

func (f *FGenre) Set(v string) {
	f.value = []string{v}
}

func (f *FGenre) Add(v string) {
	f.value = append(f.value, v)
}

func (f *FGenre) Value() string {
	if len(f.value) != 0 {
		return strings.Join(f.value, ",")
	}
	return ""
}
