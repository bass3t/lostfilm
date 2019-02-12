package channel

import (
	"strings"
)

const (
// Channel - "c"
)

type FChannel struct {
	value []string
}

func (f *FChannel) Clear() {
	f.value = []string{}
}

func (f *FChannel) Set(v string) {
	f.value = []string{v}
}

func (f *FChannel) Add(v string) {
	f.value = append(f.value, v)
}

func (f *FChannel) Value() string {
	if len(f.value) != 0 {
		return strings.Join(f.value, ",")
	}
	return ""
}
