package league

import (
	"encoding/json"
	"fmt"
)

func (l *LeagueAPI) GetSummonerByName(region, name string) (*SummonerDTO, error) {
	var summoner SummonerDTO

	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/summoner/v4/summoners/by-name/%s", region, name)

	body, err := l.DoRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &summoner)

	if err != nil {
		return nil, err
	}

	return &summoner, nil
}
