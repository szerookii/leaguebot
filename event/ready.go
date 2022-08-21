package event

import (
	"log"

	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/gateway"
	"github.com/szerookii/leaguebot/command"
	"github.com/szerookii/leaguebot/config"
)

func OnReady(client *gateway.Session, config *config.Config, cmdMgr *command.CommandManager) func() {
	return func() {
		log.Printf("Logged in as %s\n", client.Me().Tag())

		cmdMgr.Init()

		_ = client.SetActivity(&discord.Activity{Name: "les silvers feed 👺 | /help", Type: discord.ActivityListening})
	}
}
