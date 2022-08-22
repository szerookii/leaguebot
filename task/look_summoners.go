package task

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/Goscord/goscord/discord/embed"
	"github.com/szerookii/leaguebot/database"
	"github.com/szerookii/leaguebot/database/models"
)

type LookSummonersTask struct{}

func (task *LookSummonersTask) Name() string {
	return "look_summoners"
}

func (task *LookSummonersTask) GetInterval() time.Duration {
	return 30 * time.Second
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

	e := embed.NewEmbedBuilder()

	e.SetFooter(ctx.client.Me().Username, ctx.client.Me().AvatarURL())
	e.SetColor(0xe7a854)

	var dataChanged bool

	for _, oldSummonerData := range summoners {
		newSummonerData, err := ctx.leagueApi.GetSummonerByName(oldSummonerData.Region, oldSummonerData.Name)
		if err != nil {
			return err
		}

		iconUrl, _ := ctx.leagueApi.GetIconURL(newSummonerData.ProfileIconId)

		e.SetAuthor(fmt.Sprintf("%s#%s", oldSummonerData.Name, strings.ToUpper(oldSummonerData.Region)), iconUrl)
		e.SetThumbnail(iconUrl)

		if oldSummonerData.Name != newSummonerData.Name {
			e.AddField(fmt.Sprintf("Il a changé son pseudo, il est maintenant : %s", newSummonerData.Name), "C'est toujours aussi moche mais c'est pas grave...", false)
			dataChanged = true
		}

		if oldSummonerData.ProfileIconId != newSummonerData.ProfileIconId {
			e.AddField("Icône changé 😴", fmt.Sprintf("Il est peut être niveau %d mais il a de mauvais goûts", newSummonerData.SummonerLevel), false)
			dataChanged = true
		}

		if oldSummonerData.SummonerLevel != newSummonerData.SummonerLevel {
			e.AddField("Un niveau de plus, il s'agirait de toucher de l'herbe", fmt.Sprintf("Bravo %s, t'as abusé de l'EXP Boost... tu es maintenant niveau %d !", newSummonerData.Name, newSummonerData.SummonerLevel), false)
			dataChanged = true
		}

		// ToDo : Ranked stats

		if dataChanged {
			ctx.client.Channel.SendMessage(logChannel, e.Embed())

			err = db.UpdateColumns(&models.Summoner{Id: newSummonerData.Id, ProfileIconId: newSummonerData.ProfileIconId, SummonerLevel: oldSummonerData.SummonerLevel}).Error
			if err != nil {
				return err
			}
		}
	}

	return nil
}
