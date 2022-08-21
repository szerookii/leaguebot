package league

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type LeagueAPI struct {
	apiKey string
}

func NewLeagueAPI(apiKey string) *LeagueAPI {
	return &LeagueAPI{
		apiKey: apiKey,
	}
}

func (l *LeagueAPI) GetLeagueDataBySummoner(region, summonerId string) ([]*LeagueEntryDTO, error) {
	url := fmt.Sprintf("https://%s.api.riotgames.com/lol/league/v4/entries/by-summoner/%s", region, summonerId)

	body, err := l.DoRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	var rankedData []*LeagueEntryDTO

	err = json.Unmarshal(body, &rankedData)

	if err != nil {
		return nil, err
	}

	return rankedData, nil
}

func (l *LeagueAPI) DoRequest(method, url string, data io.Reader) ([]byte, error) {
	method = strings.ToUpper(method)
	req, err := http.NewRequest(method, url, data)

	if err != nil {
		return nil, err
	}

	req.Header.Set("X-Riot-Token", l.apiKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	if strings.HasPrefix(resp.Status, "4") || strings.HasPrefix(resp.Status, "5") {
		return nil, fmt.Errorf("API error: %s", resp.Status)
	}

	defer resp.Body.Close()

	var body []byte

	body, err = io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
