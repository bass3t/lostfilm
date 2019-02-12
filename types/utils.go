package types

// GetSeason return season with number num
func GetSeason(seasons []Season, num int) (s Season, ok bool) {
	for _, s := range seasons {
		if s.N == num {
			return s, true
		}
	}
	return Season{}, false
}

// GetEpisode return episode with number num from season s
func GetEpisode(s Season, num int) (e Episode, ok bool) {
	for _, e := range s.Episodes {
		if e.EpisodeNumber == num {
			return e, true
		}
	}
	return Episode{}, false
}
