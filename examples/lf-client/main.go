package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/bass3t/lostfilm"
	"github.com/bass3t/lostfilm/types"
)

func main() {

	login, pass := "user_login", "user_password"
	// serial alias, season number, episode number
	sAlias, sNum, eNum := "serial_alias", 1, 5

	lf, err := lostfilm.NewClient()
	if err != nil {
		panic(err)
	}

	captchaCB := func(captcha []byte) string {
		if len(captcha) > 0 {
			ioutil.WriteFile("captcha.gif", captcha, 0644)
		}

		fmt.Printf("Enter captcha value: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		return text
	}

	if err := lf.Login(login, pass, captchaCB); err != nil {
		panic(err)
	}

	serial, _ := lf.GetSerialByAlias(sAlias)
	seasons, _ := lf.GetSerialSeasons(serial)

	s, _ := types.GetSeason(seasons, sNum)
	e, _ := types.GetEpisode(s, eNum)

	links, _ := lf.GetEpisodeLinks(e)
	fmt.Println(links)

	return
}
