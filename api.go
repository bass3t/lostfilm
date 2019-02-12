package lostfilm

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/pkg/errors"

	"github.com/bass3t/lostfilm/filter"
	"github.com/bass3t/lostfilm/types"
	"github.com/bass3t/lostfilm/web"
)

type serialsResponse struct {
	Data   []types.Serial `json:"data"`
	Result string         `json:"result"`
}

// GetAllSerials return information about all serials
func (l *Lostfilm) GetAllSerials() ([]types.Serial, error) {
	return l.GetSerials(filter.Create())
}

// GetSerials return all filtered serials
func (l *Lostfilm) GetSerials(f filter.Filter) ([]types.Serial, error) {
	params := url.Values{}
	params.Add("act", "serial")
	params.Add("type", "search")
	params.Add("t", f.Type.Value())
	params.Add("s", f.Sort.Value())

	if f.Channel.Value() != "" {
		params.Add("c", f.Channel.Value())
	}
	if f.Genre.Value() != "" {
		params.Add("g", f.Genre.Value())
	}
	if f.Year.Value() != "" {
		params.Add("y", f.Year.Value())
	}
	if f.Group.Value() != "" {
		params.Add("r", f.Group.Value())
	}
	if f.Letter.Value() != "" {
		params.Add("l", f.Letter.Value())
	}

	var result []types.Serial
	var offset int64
	for {
		resp, err := l.sendRequest("POST", "https://www.lostfilm.tv/ajaxik.php?"+params.Encode())
		if err != nil {
			return nil, errors.Wrap(err, "request failed")
		}
		p, err := l.recvResponse(resp)
		resp.Body.Close()

		d := serialsResponse{}
		err = json.Unmarshal([]byte(p), &d)
		if err != nil {
			return nil, errors.Wrap(err, "parsing response failed")
		}
		if len(d.Data) == 0 {
			break
		}
		result = append(result, d.Data...)
		offset += 10
		params.Set("o", strconv.FormatInt(offset, 10))
	}
	return result, nil
}

// GetSerialIDByAlias return ID of serial by alias
func (l *Lostfilm) GetSerialIDByAlias(alias string) (string, error) {
	resp, err := l.sendRequest("GET", "https://www.lostfilm.tv/series/"+alias)
	if err != nil {
		return "", errors.Wrap(err, "request failed")
	}
	defer resp.Body.Close()

	p, err := l.recvResponse(resp)
	if err != nil {
		return "", errors.Wrap(err, "receiving response failed")
	}

	return web.SerialID(p)
}

// GetSerialByAlias return serial information by alias
func (l *Lostfilm) GetSerialByAlias(alias string) (s types.Serial, ok bool) {
	f := filter.Create()

	letter := string([]rune(strings.ToUpper(alias))[0])
	f.Letter.SetLetter(letter)

	checker := &types.AliasChecker{Alias: alias}
	if serials, err := l.GetSerials(f); err == nil {
		s, ok = types.FindSerial(serials, checker)
		if !ok {
			// Try find in all serial
			if serials, err := l.GetAllSerials(); err == nil {
				s, ok = types.FindSerial(serials, checker)
			}
		}
	}
	return
}

// GetSerialSeasons return all season information about serial
func (l *Lostfilm) GetSerialSeasons(s types.Serial) (seasons []types.Season, err error) {
	return l.getSerialSeasons(s.ID, s.Alias)
}

// GetSerialSeasonsByAlias return all season information about serial using serial name alias
func (l *Lostfilm) GetSerialSeasonsByAlias(alias string) (seasons []types.Season, err error) {
	id, err := l.GetSerialIDByAlias(alias)
	if err != nil {
		return seasons, errors.Wrap(err, "get serial id failed")
	}

	return l.getSerialSeasons(id, alias)
}

func (l *Lostfilm) getSerialSeasons(id, alias string) (seasons []types.Season, err error) {
	resp, err := l.sendRequest("POST", "https://www.lostfilm.tv/series/"+alias+"/seasons")
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}
	defer resp.Body.Close()

	p, err := l.recvResponse(resp)
	if err != nil {
		return nil, errors.Wrap(err, "receiving response failed")
	}

	return web.SerialSeasons(id, p)
}

// GetEpisodeLinks return links to torrent files for episode
func (l *Lostfilm) GetEpisodeLinks(e types.Episode) (links []types.EpisodeLink, err error) {
	if e.Available == false {
		return links, errors.New("can't get links of unavailable episode")
	}

	resp, err := l.sendRequest(
		"POST",
		fmt.Sprintf("https://www.lostfilm.tv/v_search.php?c=%s&s=%d&e=%d", e.SerialID, e.SeasonNumber, e.EpisodeNumber))
	if err != nil {
		return links, errors.Wrap(err, "request failed")
	}
	defer resp.Body.Close()

	p, err := l.recvResponse(resp)
	if err != nil {
		return links, errors.Wrap(err, "receiving response failed")
	}

	retreLink, err := web.GetReTreLink(p)
	if err != nil {
		return links, errors.Wrap(err, "get ReTre link failed")
	}

	resp, err = l.sendRequest("GET", retreLink)
	if err != nil {
		return links, errors.Wrap(err, "request failed")
	}
	defer resp.Body.Close()

	p, err = l.recvResponse(resp)
	if err != nil {
		return links, errors.Wrap(err, "receiving response failed")
	}

	return web.EpisodeLinks(p)
}
