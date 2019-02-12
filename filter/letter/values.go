package letter

const (
	// Letter - "l"

	A = "51"
	B = "52"
	C = "53"
	D = "54"
	E = "55"
	F = "56"
	G = "57"
	H = "58"
	I = "59"
	J = "60"
	K = "61"
	L = "62"
	M = "63"
	N = "64"
	O = "65"
	P = "66"
	Q = "67"
	R = "68"
	S = "69"
	T = "70"
	U = "71"
	V = "72"
	W = "73"
	X = "74"
	Y = "75"
	Z = "76"
)

type FLetter struct {
	value string
}

func (f *FLetter) Clear() {
	f.value = ""
}

func (f *FLetter) Set(v string) {
	f.value = v
}

func (f *FLetter) Add(v string) {
	f.value = v
}

func (f *FLetter) SetLetter(v string) {
	f.value = toFilter(v)
}

func (f *FLetter) AddLetter(v string) {
	f.value = toFilter(v)
}

func (f *FLetter) Value() string {
	return f.value
}

var fltLetter = map[string]string{
	"A": A, "B": B, "C": C, "D": D,
	"E": E, "F": F, "G": G, "H": H,
	"I": I, "J": J, "K": K, "L": L,
	"M": M, "N": N, "O": O, "P": P,
	"Q": Q, "R": R, "S": S, "T": T,
	"U": U, "V": V, "W": W, "X": X,
	"Y": Y, "Z": Z,
}

func toFilter(l string) string {
	if flt, ok := fltLetter[l]; ok {
		return flt
	}

	return ""
}
