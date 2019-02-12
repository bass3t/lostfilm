package types

// Serial contains information about serial
type Serial struct {
	Title     string `json:"title"`
	TitleOrig string `json:"title_orig"`
	Alias     string `json:"alias"`
	Channels  string `json:"channels"`
	Year      int    `json:"date,string"`
	Genres    string `json:"genres"`
	ID        string `json:"id"`
}

// Episode contains information about one episode of serial
type Episode struct {
	SerialID      string
	SeasonNumber  int
	EpisodeNumber int
	Title         string
	EngTitle      string
	Available     bool
}

// Season contains information about one season of serial
type Season struct {
	N        int
	Episodes []Episode
}

// EpisodeLink contains direct link to download torrent file of episode
type EpisodeLink struct {
	Format      string
	Quality     string
	Size        string
	TorrentLink string
}
