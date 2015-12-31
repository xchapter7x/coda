package coda

import (
	"time"

	"github.com/pivotal-pez/cfmgo"
	"gopkg.in/mgo.v2/bson"
)

type (
	//Task - a task object
	Task struct {
		ID         bson.ObjectId          `bson:"_id"`
		Timestamp  time.Time              `bson:"timestamp"`
		Expires    int64                  `bson:"expires"`
		Status     string                 `bson:"status"`
		Profile    ProfileType            `bson:"profile"`
		CallerName string                 `bson:"caller_name"`
		MetaData   map[string]interface{} `bson:"metadata"`
	}

	//TaskManagerInterface ---
	TaskManagerInterface interface {
		NewTask(callerName string, profile ProfileType, status string) (t Task)
		FindTask(id string) (t Task, err error)
		FindAndStallTaskForCaller(callerName string) (task Task, err error)
		SaveTask(t Task) (Task, error)
		ScheduleTask(t Task, expireTime time.Time) Task
	}

	//TaskSaver - interface that can save a task
	TaskSaver interface {
		SaveTask(t Task) (Task, error)
	}

	//TaskManager - manages task interactions crud stuff
	TaskManager struct {
		taskCollection cfmgo.Collection
	}

	//Agent an object which knows how to handle long running tasks. polling,
	//timeouts etc
	Agent struct {
		killTaskPoller  chan bool
		processComplete chan bool
		taskPollEmitter chan bool
		statusEmitter   chan string
		task            Task
		taskManager     TaskManagerInterface
	}

	//ProfileType - indicator of the purpose of the task to be performed
	ProfileType string
)
