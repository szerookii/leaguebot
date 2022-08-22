package task

import (
	"log"
	"time"

	"github.com/Goscord/goscord/gateway"
	"github.com/szerookii/leaguebot/config"
	"github.com/szerookii/leaguebot/league"
)

type Context struct {
	client    *gateway.Session
	config    *config.Config
	leagueApi *league.LeagueAPI
}

type Task interface {
	Name() string
	GetInterval() time.Duration
	Run(*Context) error
}

type TaskManager struct {
	client    *gateway.Session
	config    *config.Config
	leagueApi *league.LeagueAPI

	tasks map[string]Task
}

func NewTaskManager(client *gateway.Session, config *config.Config, leagueApi *league.LeagueAPI) *TaskManager {
	return &TaskManager{
		client:    client,
		leagueApi: leagueApi,
		config:    config,

		tasks: make(map[string]Task),
	}
}

func (mgr *TaskManager) Init() {
	mgr.Register(new(LookSummonersTask))
}

func (mgr *TaskManager) Register(task Task) {
	ticker := time.NewTicker(task.GetInterval())

	go func() {
		<-ticker.C
		err := task.Run(&Context{client: mgr.client, config: mgr.config, leagueApi: mgr.leagueApi})
		if err != nil {
			log.Println(err)
		}
	}()

	mgr.tasks[task.Name()] = task
}

func (mgr *TaskManager) Get(name string) Task {
	if task, ok := mgr.tasks[name]; ok {
		return task
	}

	return nil
}
