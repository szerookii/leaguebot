package league

import (
	"encoding/json"
	"fmt"
)

var regions = map[string]string{
	"euw1": "europe",
	"eun1": "europe",
	"tr1":  "europe",
	"ru":   "europe",
	"na1":  "americas",
	"br1":  "americas",
	"la1":  "americas",
	"la2":  "americas",
	"kr":   "asia",
	"jp1":  "asia",
	"oc1":  "sea",
}

func (l *LeagueAPI) GetMatchesBySummoner(region, puuid string, number int) ([]*MatchDTO, error) {
	var matchsIds []string

	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/match/v5/matches/by-puuid/%s/ids?type=ranked&count=%d", regions[region], puuid, number)

	body, err := l.DoRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &matchsIds)

	if err != nil {
		return nil, err
	}

	var matchs []*MatchDTO

	for _, matchId := range matchsIds {
		var match MatchDTO

		url := fmt.Sprintf("https://%s.api.riotgames.com/lol/match/v5/matches/%s", regions[region], matchId)

		body, err := l.DoRequest("GET", url, nil)

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(body, &match)

		if err != nil {
			return nil, err
		}

		matchs = append(matchs, &match)
	}

	return matchs, nil
}
