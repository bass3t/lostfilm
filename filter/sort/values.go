package sort

const (
	// Sort - "s"

	Rating  = "1"
	Alphbet = "2"
	New     = "3"
)

type FSort struct {
	value string
}

func (f *FSort) Clear() {
	f.value = New
}

func (f *FSort) Set(v string) {
	f.value = v
}

func (f *FSort) Add(v string) {
	f.value = v
}

func (f *FSort) Value() string {
	return f.value
}
