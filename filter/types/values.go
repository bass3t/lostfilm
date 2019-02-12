package types

const (
	// Type - "t"

	All      = "0"
	New      = "1"
	Current  = "2"
	Finished = "3"
)

type FType struct {
	value string
}

func (f *FType) Clear() {
	f.value = All
}

func (f *FType) Set(v string) {
	f.value = v
}

func (f *FType) Add(v string) {
	f.value = v
}

func (f *FType) Value() string {
	return f.value
}
