package web

import (
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/pkg/errors"

	"github.com/bass3t/lostfilm/types"
)

// SerialID parse input page and return serial identifier
func SerialID(page string) (string, error) {
	var serialID string
	rSerialID := regexp.MustCompile(`SerialId\s*=\s*([0-9]+);`)
	loc := rSerialID.FindSubmatch([]byte(page))
	if len(loc) > 1 {
		serialID = string(loc[1])
	} else {
		return "", errors.New("id not found")
	}

	return serialID, nil
}

// SerialSeasons parse input page and return information about all seasons for serial
func SerialSeasons(serialID, page string) (seasons []types.Season, err error) {
	d, err := goquery.NewDocumentFromReader(strings.NewReader(page))
	if err != nil {
		return nil, err
	}
	serieBlocks := d.Find(".series-block .serie-block")
	serieBlocks.Each(func(_ int, serieBlock *goquery.Selection) {
		d, ok := serieBlock.Find("table").Attr("id")
		if !ok {
			log.Printf("Can't find season table")
			return
		}

		if !strings.HasPrefix(d, "season_series") {
			log.Printf("Bad table id field: %s", d)
			return
		}

		sSeasonN := strings.TrimPrefix(d, "season_series_")
		seasonN, err := strconv.ParseInt(strings.TrimLeft(sSeasonN[3:6], "0"), 10, 32)
		if err != nil {
			log.Printf("Can't parse series season number(%v): %v", sSeasonN, err)
			return
		}

		if episodes := seasonEpisodes(serialID, serieBlock); len(episodes) > 0 {
			seasons = append(seasons,
				types.Season{
					N:        int(seasonN),
					Episodes: episodes,
				})
		}
	})
	return seasons, nil
}

func seasonEpisodes(serialID string, selection *goquery.Selection) (episodes []types.Episode) {
	selection.Find("table tr").Each(func(_ int, s *goquery.Selection) {
		e := types.Episode{
			SerialID: serialID,
		}
		e.Available = !s.HasClass("not-available")

		epID, ok := s.Find(".alpha .haveseen-btn").Attr("data-episode")
		if !ok {
			log.Printf("Can't parse episode number: %s", epID)
			return
		}

		if serialID != epID[:3] {
			log.Printf("Incorrect serial if need %v (recv %v)", serialID, epID[:3])
			return
		}

		num, err := strconv.ParseInt(strings.TrimLeft(epID[3:6], "0"), 10, 32)
		if err != nil {
			log.Printf("Can't parse season number(%v): %v", epID[3:6], err)
			return
		}
		e.SeasonNumber = int(num)

		num, err = strconv.ParseInt(strings.TrimLeft(epID[6:9], "0"), 10, 32)
		if err != nil {
			log.Printf("Can't parse episode number(%v): %v", epID[6:9], err)
			return
		}
		e.EpisodeNumber = int(num)

		e.Title = strings.Split(
			strings.TrimSpace(
				s.Find(".gamma").Contents().Not("span").Text()),
			"\n")[0]

		e.EngTitle = strings.TrimSpace(s.Find(".gamma span").Text())

		episodes = append(episodes, e)

	})
	return episodes
}

// GetReTreLink parse input page and return episode redirection link
func GetReTreLink(page string) (string, error) {
	var episodeLinksRegex = regexp.MustCompile("location.replace\\(\"(http.//retre.org.+)\"\\)")

	retreLinkM := episodeLinksRegex.FindAllStringSubmatch(page, -1)
	if len(retreLinkM) == 0 {
		return "", errors.New("Can't parse replace ReTre link from location")
	}
	retreLink := retreLinkM[0][1]
	return retreLink, nil
}

// EpisodeLinks parse input page and return links for episode different format torrent files
func EpisodeLinks(page string) (links []types.EpisodeLink, err error) {
	var linkDescriptionRegex = regexp.MustCompile("^Видео: (.+). Размер: (.+). Перевод: (.*)$")

	d, err := goquery.NewDocumentFromReader(strings.NewReader(page))
	if err != nil {
		return links, err
	}

	d.Find(".inner-box--list").Find(".inner-box--item").Each(func(_ int, s *goquery.Selection) {
		l := types.EpisodeLink{}
		l.Format = strings.TrimSpace(s.Find(".inner-box--label").Text())
		link, ok := s.Find(".main a").Attr("href")
		if !ok {
			log.Println("Can't find link href attribute")
			return
		}
		descr := s.Find(".inner-box--desc").Text()
		descrM := linkDescriptionRegex.FindAllStringSubmatch(descr, -1)
		if len(descrM) == 0 {
			log.Println("Can't parse link description: " + descr)
			return
		}
		l.Quality = descrM[0][1]
		l.Size = descrM[0][2]
		l.TorrentLink = link
		links = append(links, l)
	})
	return links, nil
}
