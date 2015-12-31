package coda

import (
	"fmt"
	"time"
)

//NewAgent -- creates a new initialized agent object
func NewAgent(t TaskManagerInterface, callerName string) Agent {
	return Agent{
		killTaskPoller:  make(chan bool, 1),
		processComplete: make(chan bool, 1),
		taskPollEmitter: make(chan bool, 1),
		statusEmitter:   make(chan string, 1),
		taskManager:     t,
		task:            t.NewTask(callerName, TaskAgentLongRunning, AgentTaskStatusInitializing),
	}
}

//Run - this begins the running of an agent's async process
func (s Agent) Run(process func(Agent) error) {
	s.task.Status = AgentTaskStatusRunning
	s.task.Save(s.taskManager)
	s.statusEmitter <- AgentTaskStatusRunning
	go s.startTaskPoller()
	go s.listenForPoll()
	go func(agent Agent) {
		agent.processExitHanderlDecorate(process)
		<-agent.processComplete
	}(s)
}

func (s Agent) processExitHanderlDecorate(process func(Agent) error) {
	err := process(s)
	fmt.Println("Done Agent process", process, err)
	s.taskPollEmitter <- false
	status := AgentTaskStatusComplete

	if err != nil {
		status = fmt.Sprintf("status: %s, error: %s", AgentTaskStatusFailed, err.Error())
	}
	s.task.Status = status
	s.task.Expires = 0
	s.task.Save(s.taskManager)
	s.statusEmitter <- status
}

//GetTask - get the agents task object
func (s Agent) GetTask() Task {
	return s.task
}

//GetStatus - returns a status emitting channel
func (s Agent) GetStatus() chan string {
	return s.statusEmitter
}

func (s Agent) startTaskPoller() {
ForLoop:
	for {
		select {
		case <-s.killTaskPoller:
			s.processComplete <- true
			break ForLoop
		default:
			s.taskPollEmitter <- true
		}
		time.Sleep(AgentTaskPollerInterval)
	}
}

func (s Agent) listenForPoll() {
	for <-s.taskPollEmitter {
		s.task.Expires = time.Now().Add(AgentTaskPollerTimeout).UnixNano()
		s.task.Save(s.taskManager)
	}
	s.killTaskPoller <- true
}
