package league

type SummonerDTO struct {
	Id            string `json:"id"`
	AccountId     string `json:"accountId"`
	ProfileIconId int    `json:"profileIconId"`
	RevisionDate  int64  `json:"revisionDate"`
	Name          string `json:"name"`
	Puuid         string `json:"puuid"`
	SummonerLevel int64  `json:"summonerLevel"`
}

type LeagueEntryDTO struct {
	LeagueID     string         `json:"leagueId"`
	Summoner     string         `json:"summonerId"`
	SummonerName string         `json:"summonerName"`
	QueueType    string         `json:"queueType"`
	Tier         string         `json:"tier"`
	Rank         string         `json:"rank"`
	LeaguePoints int            `json:"leaguePoints"`
	Wins         int            `json:"wins"`
	Losses       int            `json:"losses"`
	HotStreak    bool           `json:"hotStreak"`
	Veteran      bool           `json:"veteran"`
	FreshBlood   bool           `json:"freshBlood"`
	Inactive     bool           `json:"inactive"`
	MiniSeries   *MiniSeriesDTO `json:"miniSeries"`
}

type MiniSeriesDTO struct {
	Losses   int    `json:"losses"`
	Progress string `json:"progress"`
	Target   int    `json:"target"`
	Wins     int    `json:"wins"`
}

type ChampionMasteryDTO struct {
	ChampionPointsUntilNextLevel int    `json:"championPointsUntilNextLevel"`
	ChestGranted                 bool   `json:"chestGranted"`
	ChampionId                   int64  `json:"championId"`
	LastPlayTime                 int64  `json:"lastPlayTime"`
	ChampionLevel                int    `json:"championLevel"`
	SummonerId                   string `json:"summonerId"`
	ChampionPoints               int    `json:"championPoints"`
	ChampionPointsSinceLastLevel int64  `json:"championPointsSinceLastLevel"`
	TokensEarned                 int    `json:"tokensEarned"`
}

type ChampionData struct {
	Version string `json:"version"`
	Id      string `json:"id"`
	Key     int64  `json:"key,string"`
	Name    string `json:"name"`
	Title   string `json:"title"`
	Blurb   string `json:"blurb"`
	// ToDo other fields
}
