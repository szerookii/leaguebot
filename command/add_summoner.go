package command

import (
	"fmt"

	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/discord/embed"
	"github.com/szerookii/leaguebot/database"
	"github.com/szerookii/leaguebot/database/models"
	"github.com/szerookii/leaguebot/utils"
)

type AddSummonerCommand struct{}

func (c *AddSummonerCommand) Name() string {
	return "add"
}

func (c *AddSummonerCommand) Description() string {
	return "Commande admin"
}

func (c *AddSummonerCommand) Options() []*discord.ApplicationCommandOption {
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

func (c *AddSummonerCommand) Execute(ctx *Context) bool {
	if !utils.ArrayContains(ctx.config.Admins, ctx.interaction.Member.User.Id) {
		ctx.client.Interaction.CreateResponse(ctx.interaction.Id, ctx.interaction.Token, "❌ | Vous n'avez pas la permission d'utiliser cette commande.")
		return true
	}

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

	db, err := database.GetDB()
	if err != nil {
		ctx.client.Interaction.CreateResponse(ctx.interaction.Id, ctx.interaction.Token, "❌ | Impossible de se connecter à la base de données.")
		return true
	}

	summonerData := models.Summoner{
		Id:            summoner.Id,
		ProfileIconId: summoner.ProfileIconId,
		Name:          summoner.Name,
		SummonerLevel: summoner.SummonerLevel,
		Region:        region,
		Puuid:         summoner.Puuid,
	}

	for _, league := range leagues {
		if league.QueueType == "RANKED_SOLO_5x5" {
			summonerData.LastSoloTier = league.Tier
			summonerData.LastSoloRank = league.Rank
			summonerData.LastSoloLP = league.LeaguePoints
		} else if league.QueueType == "RANKED_FLEX_SR" {
			summonerData.LastFlexTier = league.Tier
			summonerData.LastFlexRank = league.Rank
			summonerData.LastFlexLP = league.LeaguePoints
		}
	}

	err = db.Create(&summonerData).Error
	if err != nil {
		ctx.client.Interaction.CreateResponse(ctx.interaction.Id, ctx.interaction.Token, "❌ | Impossible de créer le joueur dans la base de données.")
		return true
	}

	e := embed.NewEmbedBuilder()

	e.SetTitle(fmt.Sprintf("✅ %s a été ajouté dans la base de données", summoner.Name))
	e.SetDescription("Encore un silver qui va feed ?")

	e.SetThumbnail(iconUrl)
	e.SetFooter(ctx.client.Me().Username, ctx.client.Me().AvatarURL())
	e.SetColor(0xe7a854)

	ctx.client.Interaction.CreateResponse(ctx.interaction.Id, ctx.interaction.Token, e.Embed())

	return true
}
