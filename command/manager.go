package command

import (
	"log"

	"github.com/Goscord/goscord/discord"
	"github.com/Goscord/goscord/gateway"
	"github.com/szerookii/leaguebot/config"
	"github.com/szerookii/leaguebot/league"
)

type CommandManager struct {
	client    *gateway.Session
	config    *config.Config
	leagueApi *league.LeagueAPI

	commands map[string]Command
}

func NewCommandManager(client *gateway.Session, config *config.Config, leagueApi *league.LeagueAPI) *CommandManager {
	return &CommandManager{
		client:    client,
		leagueApi: leagueApi,
		config:    config,

		commands: make(map[string]Command),
	}
}

func (mgr *CommandManager) Init() {
	mgr.Register(new(HelpCommand))
	mgr.Register(new(ProfileCommand))
	mgr.Register(new(ChampionMasteryCommand))
	mgr.Register(new(AddSummonerCommand))
	mgr.Register(new(RouletteCommand))
	mgr.Register(new(InfoCommand))
}

func (mgr *CommandManager) Handler(client *gateway.Session, config *config.Config) func(*discord.Interaction) {
	return func(interaction *discord.Interaction) {
		if interaction.Member == nil {
			return
		}

		if interaction.Member.User.Bot {
			return
		}

		cmd := mgr.Get(interaction.Data.Name)

		if cmd != nil {
			go cmd.Execute(&Context{client: client, config: config, leagueApi: mgr.leagueApi, interaction: interaction, cmdMgr: mgr})
		}
	}
}

func (mgr *CommandManager) Get(name string) Command {
	if cmd, ok := mgr.commands[name]; ok {
		return cmd
	}

	return nil
}

func (mgr *CommandManager) Register(cmd Command) {
	appCmd := &discord.ApplicationCommand{
		Name:        cmd.Name(),
		Type:        discord.ApplicationCommandChat,
		Description: cmd.Description(),
		Options:     cmd.Options(),
	}

	log.Printf("Registering command %s\n", cmd.Name())
	_, err := mgr.client.Application.RegisterCommand(mgr.client.Me().Id, "", appCmd)

	if err != nil {
		log.Println(err)
	}

	mgr.commands[cmd.Name()] = cmd
}
