package main

import (
	"github.com/Goscord/goscord"
	"github.com/Goscord/goscord/gateway"
	"github.com/szerookii/leaguebot/command"
	"github.com/szerookii/leaguebot/config"
	"github.com/szerookii/leaguebot/database"
	"github.com/szerookii/leaguebot/event"
	"github.com/szerookii/leaguebot/league"
)

var (
	client    *gateway.Session
	Config    *config.Config
	leagueApi *league.LeagueAPI
	cmdMgr    *command.CommandManager
)

func main() {
	Config, _ = config.GetConfig()
	leagueApi = league.NewLeagueAPI(Config.LeagueAPIKey)
	client = goscord.New(&gateway.Options{
		Token:   Config.Token,
		Intents: gateway.IntentGuilds,
	})
	cmdMgr = command.NewCommandManager(client, Config, leagueApi)

	database.Init()

	_ = client.On("ready", event.OnReady(client, Config, cmdMgr))
	_ = client.On("interactionCreate", cmdMgr.Handler(client, Config))

	if err := client.Login(); err != nil {
		panic(err)
	}

	select {}
}
