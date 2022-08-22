package task

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Goscord/goscord/discord/embed"
	"github.com/szerookii/leaguebot/database"
	"github.com/szerookii/leaguebot/database/models"
	"github.com/szerookii/leaguebot/league"
)

var tierEmote = map[string]string{
	"IRON":        "<:iron:1011044242879164527>",
	"BRONZE":      "<:bronze:1011044168866476072>",
	"SILVER":      "<:silver:1011044248952524840>",
	"GOLD":        "<:gold:1011044227909689415>",
	"PLATINUM":    "<:platinium:1011044247086047242>",
	"DIAMOND":     "<:diamond:1011044225355350146>",
	"MASTER":      "<:master:1011044244712075374>",
	"GRANDMASTER": "<:grandmaster:1011044230107504750>",
	"CHALLENGER":  "<:challenger:1011044220942954556>",
}

type LookSummonersTask struct{}

func (task *LookSummonersTask) Name() string {
	return "look_summoners"
}

func (task *LookSummonersTask) GetInterval() time.Duration {
	return 2 * time.Minute
}

func (task *LookSummonersTask) Run(ctx *Context) error {
	db, err := database.GetDB()
	if err != nil {
		return err
	}

	logChannel := ctx.config.ReportChannelId
	if logChannel == "" {
		return errors.New("report_channel_id is not set")
	}

	if _, err = ctx.client.State().Channel(logChannel); err != nil {
		return err
	}

	var summoners []models.Summoner
	err = db.Find(&summoners).Error
	if err != nil {
		return err
	}

	var dataChanged bool

	for _, oldSummonerData := range summoners {
		newSummonerData, err := ctx.leagueApi.GetSummonerByName(oldSummonerData.Region, oldSummonerData.Name)
		if err != nil {
			return err
		}

		leagues, err := ctx.leagueApi.GetLeagueDataBySummoner(oldSummonerData.Region, newSummonerData.Id)
		if err != nil {
			return err
		}

		matchHistory, err := ctx.leagueApi.GetMatchesBySummoner(oldSummonerData.Region, newSummonerData.Puuid)
		if err != nil {
			return err
		}

		iconUrl, _ := ctx.leagueApi.GetIconURL(newSummonerData.ProfileIconId)

		if oldSummonerData.Name != newSummonerData.Name {
			e := embed.NewEmbedBuilder()

			e.SetFooter(ctx.client.Me().Username, ctx.client.Me().AvatarURL())
			e.SetColor(0xe7a854)
			e.SetAuthor(fmt.Sprintf("%s#%s", oldSummonerData.Name, strings.ToUpper(oldSummonerData.Region)), iconUrl)
			e.SetThumbnail(iconUrl)

			e.AddField(fmt.Sprintf("Il a changé son pseudo, il est maintenant : %s", newSummonerData.Name), "C'est toujours aussi moche mais c'est pas grave...", false)

			ctx.client.Channel.SendMessage(logChannel, e.Embed())

			dataChanged = true
		}

		if oldSummonerData.ProfileIconId != newSummonerData.ProfileIconId {
			e := embed.NewEmbedBuilder()

			e.SetFooter(ctx.client.Me().Username, ctx.client.Me().AvatarURL())
			e.SetColor(0xe7a854)
			e.SetAuthor(fmt.Sprintf("%s#%s", oldSummonerData.Name, strings.ToUpper(oldSummonerData.Region)), iconUrl)
			e.SetThumbnail(iconUrl)

			e.AddField("Icône changé 😴", fmt.Sprintf("Il est peut être niveau %d mais il a de mauvais goûts", newSummonerData.SummonerLevel), false)

			ctx.client.Channel.SendMessage(logChannel, e.Embed())

			dataChanged = true
		}

		if oldSummonerData.SummonerLevel != newSummonerData.SummonerLevel {
			e := embed.NewEmbedBuilder()

			e.SetFooter(ctx.client.Me().Username, ctx.client.Me().AvatarURL())
			e.SetColor(0xe7a854)
			e.SetAuthor(fmt.Sprintf("%s#%s", oldSummonerData.Name, strings.ToUpper(oldSummonerData.Region)), iconUrl)
			e.SetThumbnail(iconUrl)

			e.AddField("Un niveau de plus, il s'agirait de toucher de l'herbe", fmt.Sprintf("Bravo %s, t'as abusé de l'EXP Boost... tu es maintenant niveau %d !", newSummonerData.Name, newSummonerData.SummonerLevel), false)

			ctx.client.Channel.SendMessage(logChannel, e.Embed())

			dataChanged = true
		}

		var soloUpated bool = false
		var flexUpdated bool = false

		for _, match := range matchHistory {
			if match.Info.QueueId == league.RankedSolo && !soloUpated {
				soloUpated = true

				if match.Info.GameId != oldSummonerData.LastSoloGameId {
					dataChanged = true
					leagueData := getLeagueData(leagues, "RANKED_SOLO_5x5")

					e := embed.NewEmbedBuilder()

					e.SetFooter(ctx.client.Me().Username, ctx.client.Me().AvatarURL())
					e.SetAuthor(fmt.Sprintf("%s#%s", oldSummonerData.Name, strings.ToUpper(oldSummonerData.Region)), iconUrl)
					e.SetThumbnail(iconUrl)

					if wonGame(match.Info.Participants, newSummonerData.Id) {
						e.SetColor(embed.Green)

						e.AddField("Petite victoire en full tryhard (ou pas)", "Beaucoup pensent qu'en ranked on se doit de fumer tout les autres joueurs, mais non, on finit le nexus même en étant en 0/10", false)
						e.AddField("Victoire", "✅", true)
						e.AddField("Rang", fmt.Sprintf("%s %s %s", tierEmote[leagueData.Tier], leagueData.Tier, leagueData.Rank), true)
						e.AddField("LP gagnés", fmt.Sprintf(":small_red_triangle: %d", leagueData.LeaguePoints-oldSummonerData.LastSoloLP), true)
						e.AddField("LP", fmt.Sprintf(":trophy: %d", leagueData.LeaguePoints), true)
					} else {
						e.SetColor(embed.Red)

						e.AddField("De retour en looserQ", "> Qu'est ce que la honte ?\n> Sentiment d'abaissement, d'humiliation qui résulte d'une atteinte à l'honneur, à la dignité\nLa honte, c'est aussi perdre en ranked sur LoL, on espère ne plus te revoir", false)
						e.AddField("Victoire", "❌", true)
						e.AddField("Rang", fmt.Sprintf("%s %s %s", tierEmote[leagueData.Tier], leagueData.Tier, leagueData.Rank), true)
						e.AddField("LP perdus", fmt.Sprintf(":small_red_triangle_down: %d LP", oldSummonerData.LastSoloLP-leagueData.LeaguePoints), true)
						e.AddField("LP", fmt.Sprintf(":trophy: %d LP", leagueData.LeaguePoints), true)
					}

					oldSummonerData.LastSoloGameId = match.Info.GameId
					oldSummonerData.LastSoloTier = leagueData.Tier
					oldSummonerData.LastSoloRank = leagueData.Rank
					oldSummonerData.LastSoloLP = leagueData.LeaguePoints

					ctx.client.Channel.SendMessage(logChannel, e.Embed())
				}
			}

			if match.Info.QueueId == league.RankedFlex && !flexUpdated {
				flexUpdated = true

				if match.Info.GameId != oldSummonerData.LastFlexGameId {
					dataChanged = true
					leagueData := getLeagueData(leagues, "RANKED_FLEX_SR")

					e := embed.NewEmbedBuilder()

					e.SetFooter(ctx.client.Me().Username, ctx.client.Me().AvatarURL())
					e.SetColor(0xe7a854)
					e.SetAuthor(fmt.Sprintf("%s#%s", oldSummonerData.Name, strings.ToUpper(oldSummonerData.Region)), iconUrl)
					e.SetThumbnail(iconUrl)

					if wonGame(match.Info.Participants, newSummonerData.Id) {
						e.SetColor(embed.Green)

						e.AddField("Petite victoire en full chill (ou pas)", "> Qu'est ce que la classée flexible ?\n> La classée flexible dans League of Legends est un mode de jeu qui est classé séparément de la file d'attente solo/duo. Il s'agit d'un mode de jeu compétitif à cinq contre cinq où vous pouvez avoir un groupe de 1, 2, 3 ou 5 joueurs (à l'exception des groupes de 4) dans un cadre classé.\nEn gros, tu joues à 5 et tu te fais hardcarry...", false)
						e.AddField("Victoire", "✅", true)
						e.AddField("Rang", fmt.Sprintf("%s %s %s", tierEmote[leagueData.Tier], leagueData.Tier, leagueData.Rank), true)
						e.AddField("LP gagnés", fmt.Sprintf(":small_red_triangle: %d", leagueData.LeaguePoints-oldSummonerData.LastFlexLP), true)
						e.AddField("LP", fmt.Sprintf(":trophy: %d", leagueData.LeaguePoints), true)
					} else {
						e.SetColor(embed.Red)

						e.AddField("De retour en looserQ (en flexible en plus, la honte)", "> Qu'est ce que la classée flexible ?\n> La classée flexible dans League of Legends est un mode de jeu qui est classé séparément de la file d'attente solo/duo. Il s'agit d'un mode de jeu compétitif à cinq contre cinq où vous pouvez avoir un groupe de 1, 2, 3 ou 5 joueurs (à l'exception des groupes de 4) dans un cadre classé.\nEn gros, tu joues à 5 et tu te fais hardcarry, sauf que là c’était pas ton cas avec le powerspike en 0/10...", false)
						e.AddField("Victoire", "❌", true)
						e.AddField("Rang", fmt.Sprintf("%s %s %s", tierEmote[leagueData.Tier], leagueData.Tier, leagueData.Rank), true)
						e.AddField("LP perdus", fmt.Sprintf(":small_red_triangle_down: %d LP", oldSummonerData.LastFlexLP-leagueData.LeaguePoints), true)
						e.AddField("LP", fmt.Sprintf(":trophy: %d LP", leagueData.LeaguePoints), true)
					}

					oldSummonerData.LastFlexGameId = match.Info.GameId
					oldSummonerData.LastFlexTier = leagueData.Tier
					oldSummonerData.LastFlexRank = leagueData.Rank
					oldSummonerData.LastFlexLP = leagueData.LeaguePoints

					ctx.client.Channel.SendMessage(logChannel, e.Embed())
				}
			}
		}

		if dataChanged {
			err = db.UpdateColumns(&models.Summoner{
				Id:             newSummonerData.Id,
				ProfileIconId:  newSummonerData.ProfileIconId,
				SummonerLevel:  oldSummonerData.SummonerLevel,
				LastSoloGameId: oldSummonerData.LastSoloGameId,
				LastSoloTier:   oldSummonerData.LastSoloTier,
				LastSoloRank:   oldSummonerData.LastSoloRank,
				LastSoloLP:     oldSummonerData.LastSoloLP,
				LastFlexGameId: oldSummonerData.LastFlexGameId,
				LastFlexTier:   oldSummonerData.LastFlexTier,
				LastFlexRank:   oldSummonerData.LastFlexRank,
				LastFlexLP:     oldSummonerData.LastFlexLP,
			}).Error
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func wonGame(participants []*league.ParticipantDTO, summonerId string) bool {
	for _, participant := range participants {
		if participant.SummonerId == summonerId {
			return participant.Win
		}
	}

	return false
}

func getLeagueData(leagues []*league.LeagueEntryDTO, queueType string) *league.LeagueEntryDTO {
	for _, league := range leagues {
		if league.QueueType == queueType {
			return league
		}
	}

	return nil
}
