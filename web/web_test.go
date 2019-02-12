package web

import (
	"testing"

	"github.com/bass3t/lostfilm/types"
)

const serialSeasonPage = `
<div class="series-block">
<div class="serie-block">
<table cellpadding="0" cellspacing="0" width="100%" class="movie-parts-list" id="season_series_111003999">
<tbody>
<tr class="not-available">
<td class="alpha">
<div class="haveseen-btn" title="" data-episode="111003007" data-season="111003999"></div>
</td>
<td class="gamma" onClick="goTo('/series/Serial_Name/season_3/episode_7/',false)">The End <br><span class="small-text">The End </span></td>
</tr>
<tr>
<td class="alpha">
<div class="haveseen-btn" title="" onClick="markEpisodeAsWatched(this);" data-episode="111003006" data-season="111003999" data-code="111-3-6"></div>
</td>
<td class="gamma" onClick="goTo('/series/Serial_Name/season_3/episode_6/',false)" title="">						
<div>
Серия 3-6.<br />
<span class="gray-color2 small-text">The number 3-6</span>
</div>
</td>
</tr>
</tbody>
</table>
</div>
<div class="serie-block">
<table cellpadding="0" cellspacing="0" width="100%" class="movie-parts-list" id="season_series_111002999">
<tbody>
<tr>
<td class="alpha">
<div class="haveseen-btn" title="" onClick="markEpisodeAsWatched(this);" data-episode="111002010" data-season="111002999" data-code="111-2-10"></div>
</td>
<td class="gamma" onClick="goTo('/series/Serial_Name/season_2/episode_10/',false)" title="">						
<div>
Серия 2-10.<br />
<span class="gray-color2 small-text">The number 2-10</span>
</div>
</td>
</tr>
<tr>
<td class="alpha">
<div class="haveseen-btn" title="" onClick="markEpisodeAsWatched(this);" data-episode="111002009" data-season="111002999" data-code="111-2-9"></div>
</td>
<td class="gamma" onClick="goTo('/series/Serial_Name/season_2/episode_9/',false)" title="">						
<div>
Серия 2-9.<br />
<span class="gray-color2 small-text">The number 2-9</span>
</div>
</td>
</tr>
</tbody>
</table>
</div>
</div>
`

var seasonsCheck = []types.Season{
	{
		N: 3,
		Episodes: []types.Episode{
			{
				EpisodeNumber: 7,
				SerialID:      "111",
				SeasonNumber:  3,
				Title:         "The End",
				EngTitle:      "The End",
				Available:     false,
			},
			{
				EpisodeNumber: 6,
				SerialID:      "111",
				SeasonNumber:  3,
				Title:         "Серия 3-6.",
				EngTitle:      "The number 3-6",
				Available:     true,
			},
		},
	},
	{
		N: 2,
		Episodes: []types.Episode{
			{
				EpisodeNumber: 10,
				SerialID:      "111",
				SeasonNumber:  2,
				Title:         "Серия 2-10.",
				EngTitle:      "The number 2-10",
				Available:     true,
			},
			{
				EpisodeNumber: 9,
				SerialID:      "111",
				SeasonNumber:  2,
				Title:         "Серия 2-9.",
				EngTitle:      "The number 2-9",
				Available:     true,
			},
		},
	},
}

const serialIDPage = `
<div class="title-block"></div>
<div class="clr"></div>
<script type="text/javascript">
SerialId = 111;
$(function(){});
`

const serialIDCheck = "111"

const retreLinkPage = `
<head>
<meta http-equiv="refresh" content="0; url=http://retre.org/v3/index.php?c=111&s=2&e=10&u=1234567&h=00112233445566778899aabbccddeeff&n=1">
<style type="text/css">
*
{
	font-family:Tahoma;
	font-size:12px;
	color:#000000;
}
</style>
<script type="text/javascript">
function r()
<!--
location.replace("http://retre.org/v3/index.php?c=111&s=2&e=10&u=1234567&h=00112233445566778899aabbccddeeff&n=1");
//-->
</script>
</head>
`

const retreLinkCheck = "http://retre.org/v3/index.php?c=111&s=2&e=10&u=1234567&h=00112233445566778899aabbccddeeff&n=1"

const episodeLinksPage = `
<div class="inner-box--list">
<div class="inner-box--item">
<div class="inner-box--label">SD</div>
<div class="inner-box--link main"><a href="http://tracktor.in/td.php?s=link1">2 сезон, 10 серия. WEBRip</a></div>
<div class="inner-box--link sub"><a href="http://tracktor.in/td.php?s=link1">http://tracktor.in/td.php?s=link1</a></div>
<div class="inner-box--desc">Видео: WEBRip. Размер: 100 МБ. Перевод: Многоголосый закадровый</div>
</div>
<div class="inner-box--item">
<div class="inner-box--label">1080</div>
<div class="inner-box--link main"><a href="http://tracktor.in/td.php?s=link2">2 сезон, 10 серия. 1080p WEBRip</a></div>
<div class="inner-box--link sub"><a href="http://tracktor.in/td.php?s=link2">http://tracktor.in/td.php?s=link2</a></div>
<div class="inner-box--desc">Видео: 1080p WEBRip. Размер: 1.8 ГБ. Перевод: Многоголосый закадровый</div>
</div>
<div class="inner-box--item">
<div class="inner-box--label">MP4</div>
<div class="inner-box--link main"><a href="http://tracktor.in/td.php?s=link3">2 сезон, 10 серия. 720p WEBRip</a></div>
<div class="inner-box--link sub"><a href="http://tracktor.in/td.php?s=link3">http://tracktor.in/td.php?s=link3</a></div>
<div class="inner-box--desc">Видео: 720p WEBRip. Размер: 1 ГБ. Перевод: Многоголосый закадровый</div>
</div>
</div>
`

var episodeLinksCheck = []types.EpisodeLink{
	{
		Format:      "SD",
		Quality:     "WEBRip",
		Size:        "100 МБ",
		TorrentLink: "http://tracktor.in/td.php?s=link1",
	},
	{
		Format:      "1080",
		Quality:     "1080p WEBRip",
		Size:        "1.8 ГБ",
		TorrentLink: "http://tracktor.in/td.php?s=link2",
	},
	{
		Format:      "MP4",
		Quality:     "720p WEBRip",
		Size:        "1 ГБ",
		TorrentLink: "http://tracktor.in/td.php?s=link3",
	},
}

func checkEpisode(e1, e2 types.Episode) bool {
	return (e1.EpisodeNumber == e2.EpisodeNumber) &&
		(e1.SerialID == e2.SerialID) &&
		(e1.SeasonNumber == e2.SeasonNumber) &&
		(e1.Title == e2.Title) &&
		(e1.EngTitle == e2.EngTitle) &&
		(e1.Available == e2.Available)
}

func checkEpisodes(e1, e2 []types.Episode) bool {
	if len(e1) != len(e2) {
		return false
	}

	for i := 0; i < len(e1); i++ {
		if !checkEpisode(e1[i], e2[i]) {
			return false
		}
	}

	return true
}

func checkSeason(s1, s2 types.Season) bool {
	return (s1.N == s2.N) && checkEpisodes(s1.Episodes, s2.Episodes)
}

func checkSeasons(s1, s2 []types.Season) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := 0; i < len(s1); i++ {
		if checkSeason(s1[i], s2[i]) != true {
			return false
		}
	}

	return true
}

func checkEpisodeLink(l1, l2 types.EpisodeLink) bool {
	return (l1.Format == l2.Format) &&
		(l1.Quality == l2.Quality) &&
		(l1.Size == l2.Size) &&
		(l1.TorrentLink == l2.TorrentLink)
}

func checkEpisodeLinks(l1, l2 []types.EpisodeLink) bool {
	if len(l1) != len(l2) {
		return false
	}

	for i := 0; i < len(l1); i++ {
		if !checkEpisodeLink(l1[i], l2[i]) {
			return false
		}
	}

	return true
}

func TestSerialSeasons(t *testing.T) {
	seasons, err := SerialSeasons("111", serialSeasonPage)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(seasons) != 2 {
		t.Fatal("Must be 2 seasons")
	}

	if !checkSeasons(seasons, seasonsCheck) {
		t.Fatal("Incorrect seasons information")
	}
}

func TestSerialID(t *testing.T) {
	serialID, err := SerialID(serialIDPage)
	if err != nil {
		t.Fatal(err.Error())
	}

	if serialID != serialIDCheck {
		t.Fatal("Incorrect serialID value")
	}
}

func TestReTreLink(t *testing.T) {
	retreLink, err := GetReTreLink(retreLinkPage)
	if err != nil {
		t.Fatal(err.Error())
	}

	if retreLink != retreLinkCheck {
		t.Fatal("Incorrect retreLink value")
	}
}

func TestEpisodeLinks(t *testing.T) {
	links, err := EpisodeLinks(episodeLinksPage)
	if err != nil {
		t.Fatal(err.Error())
	}

	if len(links) != 3 {
		t.Fatal("Must be 3 links")
	}

	if !checkEpisodeLinks(links, episodeLinksCheck) {
		t.Fatal("Incorrect links information")
	}
}
