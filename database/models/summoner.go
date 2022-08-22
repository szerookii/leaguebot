package models

type Summoner struct {
	Id            string `json:"id"`
	ProfileIconId int    `json:"profileIconId"`
	Name          string `json:"name"`
	SummonerLevel int64  `json:"summonerLevel"`
	Puuid         string `json:"puuid"`
	Region        string `json:"region"`

	LastSoloGameId int64  `json:"lastGameId"`
	LastSoloTier   string `json:"lastTier"`
	LastSoloRank   string `json:"lastRank"`
	LastSoloLP     int    `json:"lastLP"`

	LastFlexGameId int64  `json:"lastFlexGameId"`
	LastFlexTier   string `json:"lastFlexTier"`
	LastFlexRank   string `json:"lastFlexRank"`
	LastFlexLP     int    `json:"lastFlexLP"`
}
