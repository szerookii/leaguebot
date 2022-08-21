package command

import (
	"fmt"

	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/discord/embed"
)

type HelpCommand struct{}

func (c *HelpCommand) Name() string {
	return "help"
}

func (c *HelpCommand) Description() string {
	return "Affiche la page d'aide!"
}

func (c *HelpCommand) Options() []*discord.ApplicationCommandOption {
	return make([]*discord.ApplicationCommandOption, 0)
}

func (c *HelpCommand) Execute(ctx *Context) bool {
	e := embed.NewEmbedBuilder()

	e.SetTitle(":books: | Page d'aide")

	for _, cmd := range ctx.cmdMgr.commands {
		e.AddField(fmt.Sprintf("/%s", cmd.Name()), cmd.Description(), false)
	}

	e.SetFooter(ctx.client.Me().Username, ctx.client.Me().AvatarURL())
	e.SetColor(0xe7a854)

	ctx.client.Interaction.CreateResponse(ctx.interaction.Id, ctx.interaction.Token, e.Embed())

	return true
}
