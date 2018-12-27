/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/27
 * @Last Modified by: liliangshan
 * *************************************************/

package cron

import "time"

type Scheduler struct {
	tasks        []*Task
	stop         chan struct{}
	add          chan *Task
	isRemovedJob chan IsRemovedJob
	snapshot     chan []*Task
	running      bool
}

type IsRemovedJob func(task *Task) bool

type Task struct {
	Cron Cron

	Next time.Time

	Prev time.Time

	Runner Runner
}

type Runner interface {
	Execute()
}

func New() *Scheduler {
	return &Scheduler{
		tasks:        nil,
		stop:         make(chan struct{}),
		add:          make(chan *Task),
		isRemovedJob: make(chan IsRemovedJob),
		snapshot:     make(chan []*Task),
		running:      false,
	}
}

type WrapperRunner func()

func (wrapperRunner WrapperRunner) Execute() {
	wrapperRunner()
}

func (scheduler Scheduler) SubmitTask(cron Cron, runner Runner) {
	task := &Task{
		Cron:   cron,
		Runner: runner,
	}
	if !scheduler.running {
		scheduler.tasks = append(scheduler.tasks, task)
		return
	}

	scheduler.add <- task
}

func (scheduler Scheduler) AddTask(spec string, runner Runner) error {
	return nil
}
