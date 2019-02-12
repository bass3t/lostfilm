package filter

import (
	"github.com/bass3t/lostfilm/filter/channel"
	"github.com/bass3t/lostfilm/filter/genre"
	"github.com/bass3t/lostfilm/filter/group"
	"github.com/bass3t/lostfilm/filter/letter"
	"github.com/bass3t/lostfilm/filter/sort"
	"github.com/bass3t/lostfilm/filter/types"
	"github.com/bass3t/lostfilm/filter/year"
)

// Available filteres
// Genre - "g"
// Sort - "s"
// Type - "t"
// Year - "y"
// Channel - "c"
// Group - "r"
// Letter - "l"

// Filter describe filters for reuest serials list
type Filter struct {
	Sort    sort.FSort
	Type    types.FType
	Genre   genre.FGenre
	Year    year.FYear
	Channel channel.FChannel
	Group   group.FGroup
	Letter  letter.FLetter
}

// Create new filter for request serials list
func Create() (f Filter) {
	f.Sort.Clear()
	f.Type.Clear()
	f.Genre.Clear()
	f.Year.Clear()
	f.Channel.Clear()
	f.Group.Clear()
	f.Letter.Clear()
	return
}
