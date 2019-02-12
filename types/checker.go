package types

type Checker interface {
	Check(s Serial) bool
}

type AliasChecker struct {
	Alias string
}

func (c *AliasChecker) Check(s Serial) bool {
	return s.Alias == c.Alias
}

// FindSerial in slice of serials
// for each searial in serials call function f
// return serial and true, if f return true
// return false if f return false for all serials
func FindSerial(serials []Serial, c Checker) (s Serial, ok bool) {
	for _, v := range serials {
		if c.Check(v) {
			return v, true
		}
	}

	return Serial{}, false
}
