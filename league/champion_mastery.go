package league

import (
	"encoding/json"
	"fmt"
)

func (l *LeagueAPI) GetChampionMasteriesBySummoner(region, summonerId string) ([]*ChampionMasteryDTO, error) {
	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/champion-mastery/v4/champion-masteries/by-summoner/%s", region, summonerId)

	body, err := l.DoRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	var championMasteries []*ChampionMasteryDTO
	err = json.Unmarshal(body, &championMasteries)

	if err != nil {
		return nil, err
	}

	return championMasteries, nil
}
