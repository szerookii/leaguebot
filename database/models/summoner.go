package models

type Summoner struct {
	Id            string `json:"id"`
	ProfileIconId int    `json:"profileIconId"`
	AccountName   string `json:"accountName"`
	SummonerLevel int64  `json:"summonerLevel"`
}
