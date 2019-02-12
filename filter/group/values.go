package group

import (
	"strings"
)

const (
// Group - "r"
)

type FGroup struct {
	value []string
}

func (f *FGroup) Clear() {
	f.value = []string{}
}

func (f *FGroup) Set(v string) {
	f.value = []string{v}
}

func (f *FGroup) Add(v string) {
	f.value = append(f.value, v)
}

func (f *FGroup) Value() string {
	if len(f.value) != 0 {
		return strings.Join(f.value, ",")
	}
	return ""
}
