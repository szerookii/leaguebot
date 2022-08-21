package command

import (
	"fmt"

	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/discord/embed"
	"github.com/szerookii/leaguebot/utils"
)

var masteryEmotes = []string{
	"",
	"<:mastery1:896783387732369418>",
	"<:mastery2:896783399337996368>",
	"<:mastery3:896783407126822945>",
	"<:mastery4:896783414466871337>",
	"<:mastery5:896783422134030368>",
	"<:mastery6:896783430786879498>",
	"<:mastery7:896783437581672488>",
}

type ChampionMasteryCommand struct{}

func (c *ChampionMasteryCommand) Name() string {
	return "championmastery"
}

func (c *ChampionMasteryCommand) Description() string {
	return "Afficher la maîtrise de champion du joueur!"
}

func (c *ChampionMasteryCommand) Options() []*discord.ApplicationCommandOption {
	return []*discord.ApplicationCommandOption{
		{
			Type:        discord.ApplicationCommandOptionString,
			Name:        "summoner",
			Description: "Le nom du joueur",
			Required:    true,
		},

		{
			Type:        discord.ApplicationCommandOptionString,
			Name:        "region",
			Description: "La région du joueur",
			Required:    true,
			Choices: []*discord.ApplicationCommandOptionChoice{
				{
					Name:  "EUW",
					Value: "euw1",
				},
				{
					Name:  "EUNE",
					Value: "eun1",
				},
				{
					Name:  "NA",
					Value: "na1",
				},
				{
					Name:  "OCE",
					Value: "oc1",
				},
				{
					Name:  "LAN",
					Value: "la1",
				},
				{
					Name:  "LAS",
					Value: "la2",
				},
				{
					Name:  "BR",
					Value: "br1",
				},
				{
					Name:  "RU",
					Value: "ru",
				},
				{
					Name:  "TR",
					Value: "tr1",
				},
				{
					Name:  "JP",
					Value: "jp1",
				},
			},
		},
	}
}

func (c *ChampionMasteryCommand) Execute(ctx *Context) bool {
	summonerName := ctx.interaction.Data.Options[0].String()
	region := ctx.interaction.Data.Options[1].String()

	summoner, err := ctx.leagueApi.GetSummonerByName(region, summonerName)
	if err != nil {
		ctx.client.Interaction.CreateResponse(ctx.interaction.Id, ctx.interaction.Token, "❌ | Impossible de trouver le joueur!")
		return true
	}

	masteries, err := ctx.leagueApi.GetChampionMasteriesBySummoner(region, summoner.Id)
	if err != nil {
		ctx.client.Interaction.CreateResponse(ctx.interaction.Id, ctx.interaction.Token, "❌ | Impossible de trouver le joueur!")
		return true
	}

	iconUrl, _ := ctx.leagueApi.GetIconURL(summoner.ProfileIconId)
	e := embed.NewEmbedBuilder()

	e.SetTitle(fmt.Sprintf(":trophy: | Maîtrises de %s", summoner.Name))

	for i, mastery := range masteries {
		if i > 8 {
			break
		}

		champion, err := ctx.leagueApi.GetChampionDataById("fr_FR", mastery.ChampionId)

		if err != nil {
			continue
		}

		e.AddField(fmt.Sprintf("%s %s", masteryEmotes[mastery.ChampionLevel], champion.Name), fmt.Sprintf("Niveau : %d\nPoints : %s", mastery.ChampionLevel, utils.FormatNumber(mastery.ChampionPoints)), true)
	}

	e.SetThumbnail(iconUrl)
	e.SetFooter(ctx.client.Me().Username, ctx.client.Me().AvatarURL())
	e.SetColor(0xe7a854)

	ctx.client.Interaction.CreateResponse(ctx.interaction.Id, ctx.interaction.Token, e.Embed())

	return true
}
