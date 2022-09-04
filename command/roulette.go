package command

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/Goscord/goscord/discord"
	"github.com/szerookii/leaguebot/utils"
)

type RouletteCommand struct{}

func (c *RouletteCommand) Name() string {
	return "rr"
}

func (c *RouletteCommand) Description() string {
	return "Commande admin"
}

func (c *RouletteCommand) Options() []*discord.ApplicationCommandOption {
	return []*discord.ApplicationCommandOption{
		{
			Type:        discord.ApplicationCommandOptionString,
			Name:        "choix",
			Description: "Les choix, séparés par une virgule",
			Required:    true,
		},
	}
}

func (c *RouletteCommand) Execute(ctx *Context) bool {
	if !utils.ArrayContains(ctx.config.Admins, ctx.interaction.Member.User.Id) {
		ctx.client.Interaction.CreateResponse(ctx.interaction.Id, ctx.interaction.Token, "❌ | Vous n'avez pas la permission d'utiliser cette commande.")
		return true
	}

	choices := strings.Split(ctx.interaction.Data.Options[0].String(), ",")

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(choices), func(i, j int) { choices[i], choices[j] = choices[j], choices[i] })

	m, err := ctx.client.Interaction.CreateResponse(ctx.interaction.Id, ctx.interaction.Token, "Tirage au sort...")

	if err != nil {
		return false
	}

	if m.Type != discord.InteractionCallbackTypeDeferredChannelMessageWithSource {
		return false
	}

	cId := m.Data.(*discord.Message).ChannelId
	id := m.Data.(*discord.Message).Id

	var choice string

	for i, ci := 0, 0; i < 5; i++ {
		if ci > len(choices) {
			ci = 0
		}

		choice := choices[ci]

		ctx.client.Channel.Edit(cId, id, fmt.Sprintf("➡️ %s ⬅️", choice))

		if err != nil {
			fmt.Println(err)
		}

		time.Sleep(1 * time.Second)
	}

	ctx.client.Channel.Edit(cId, id, fmt.Sprintf("Le résultat de la roulette est ➡️ %s", choice))

	return true
}
