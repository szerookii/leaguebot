package league

import (
	"encoding/json"
	"fmt"
)

func (l *LeagueAPI) GetCurrentVersion() (string, error) {
	var versions []string

	body, err := l.DoRequest("GET", "https://ddragon.leagueoflegends.com/api/versions.json", nil)

	if err != nil {
		return "", err
	}

	err = json.Unmarshal(body, &versions)

	if err != nil {
		return "", err
	}

	return versions[0], nil
}

func (l *LeagueAPI) GetIconURL(iconId int) (string, error) {
	version, err := l.GetCurrentVersion()

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("http://ddragon.leagueoflegends.com/cdn/%s/img/profileicon/%d.png", version, iconId), nil
}

type championsDataResponse struct {
	Data map[string]*ChampionData `json:"data"`
}

func (l *LeagueAPI) GetChampionDataById(lang string, championId int64) (*ChampionData, error) {
	version, err := l.GetCurrentVersion()

	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("https://ddragon.leagueoflegends.com/cdn/%s/data/%s/champion.json", version, lang)
	body, err := l.DoRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	var champions championsDataResponse
	err = json.Unmarshal(body, &champions)

	if err != nil {
		return nil, err
	}

	for _, champion := range champions.Data {
		if championId == champion.Key {
			return champion, nil
		}
	}

	return nil, nil
}
