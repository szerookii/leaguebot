package league

type QueueType int

const (
	RankedSolo QueueType = 420
	RankedFlex QueueType = 440
	Clash      QueueType = 700
)

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

type MatchDTO struct {
	Metadata *MetadataDTO `json:"metadata"`
	Info     *InfoDTO     `json:"info"`
}

type MetadataDTO struct {
	DataVersion  string   `json:"dataVersion"`
	MatchId      string   `json:"matchId"`
	Participants []string `json:"participants"` // A list of participant PUUIDs.
}

type InfoDTO struct {
	GameCreation       int64             `json:"gameCreation"`
	GameDuration       int64             `json:"gameDuration"`
	GameEndTimestamp   int64             `json:"gameEndTimestamp"`
	GameId             int64             `json:"gameId"`
	GameMode           string            `json:"gameMode"`
	GameName           string            `json:"gameName"`
	GameStartTimestamp int64             `json:"gameStartTimestamp"`
	GameType           string            `json:"gameType"`
	GameVersion        string            `json:"gameVersion"`
	MapId              int64             `json:"mapId"`
	Participants       []*ParticipantDTO `json:"participants"`
	PlatformId         string            `json:"platformId"`
	QueueId            int64             `json:"queueId"`
	//Teams 	[]*TeamDTO `json:"teams"`
	TournamentCode string `json:"tournamentCode"`
}

type ParticipantDTO struct {
	// ToDo : Add other fields

	Win bool `json:"win"`
}
