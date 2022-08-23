package command

import (
	"runtime"

	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/discord/embed"
)

type InfoCommand struct{}

func (c *InfoCommand) Name() string {
	return "info"
}

func (c *InfoCommand) Description() string {
	return "Affiche quelques informations utiles à propos du bot!"
}

func (c *InfoCommand) Options() []*discord.ApplicationCommandOption {
	return make([]*discord.ApplicationCommandOption, 0)
}

func (c *InfoCommand) Execute(ctx *Context) bool {
	var mStats runtime.MemStats
	runtime.ReadMemStats(&mStats)

	e := embed.NewEmbedBuilder()

	e.SetTitle(":interrobang: | Informations")

	e.AddField("Développeur", "<@810596177857871913>", false)
	e.AddField("Dépôt GitHub", "[Cliquez ici](https://github.com/szerookii/leaguebot)", false)
	e.AddField("API Discord", "[Goscord](https://goscord.dev)", false)

	e.SetFooter(ctx.client.Me().Username, ctx.client.Me().AvatarURL())
	e.SetColor(0xe7a854)

	ctx.client.Interaction.CreateResponse(ctx.interaction.Id, ctx.interaction.Token, e.Embed())

	return true
}
