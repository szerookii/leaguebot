package command

import (
	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/gateway"
	"github.com/szerookii/leaguebot/config"
	"github.com/szerookii/leaguebot/league"
)

type Context struct {
	cmdMgr      *CommandManager
	config      *config.Config
	leagueApi   *league.LeagueAPI
	client      *gateway.Session
	interaction *discord.Interaction
}

type Command interface {
	Name() string
	Description() string
	Options() []*discord.ApplicationCommandOption
	Execute(ctx *Context) bool
}
