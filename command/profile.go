package command

import (
	"fmt"

	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/discord/embed"
)

var queueType = map[string]string{
	"RANKED_SOLO_5x5": "Classée Solo/Duo",
	"RANKED_FLEX_SR":  "Classée Flexible",
	"RANKED_FLEX_TT":  "Classée Flexible 3v3 (TFT)",
}

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

type ProfileCommand struct{}

func (c *ProfileCommand) Name() string {
	return "profile"
}

func (c *ProfileCommand) Description() string {
	return "Affiche les informations d'un joueur!"
}

func (c *ProfileCommand) Options() []*discord.ApplicationCommandOption {
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

func (c *ProfileCommand) Execute(ctx *Context) bool {
	summonerName := ctx.interaction.Data.Options[0].String()
	region := ctx.interaction.Data.Options[1].String()

	summoner, err := ctx.leagueApi.GetSummonerByName(region, summonerName)
	if err != nil {
		ctx.client.Interaction.CreateResponse(ctx.interaction.Id, ctx.interaction.Token, "❌ | Impossible de trouver le joueur!")
		return true
	}

	leagues, err := ctx.leagueApi.GetLeagueDataBySummoner(region, summoner.Id)
	if err != nil {
		ctx.client.Interaction.CreateResponse(ctx.interaction.Id, ctx.interaction.Token, "❌ | Impossible de trouver le joueur!")
		return true
	}

	iconUrl, _ := ctx.leagueApi.GetIconURL(summoner.ProfileIconId)
	e := embed.NewEmbedBuilder()

	e.SetTitle(fmt.Sprintf(":trophy: | Profil de %s", summoner.Name))

	e.AddField("Niveau", fmt.Sprintf("%d", summoner.SummonerLevel), false)

	for _, league := range leagues {
		e.AddField(queueType[league.QueueType], fmt.Sprintf("%s %s %s %d LP (%d victoires / %d défaites)", tierEmote[league.Tier], league.Tier, league.Rank, league.LeaguePoints, league.Wins, league.Losses), false)
	}

	e.SetThumbnail(iconUrl)
	e.SetFooter(ctx.client.Me().Username, ctx.client.Me().AvatarURL())
	e.SetColor(0xe7a854)

	ctx.client.Interaction.CreateResponse(ctx.interaction.Id, ctx.interaction.Token, e.Embed())

	return true
}
