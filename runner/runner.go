/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/31
 * @Last Modified by: liliangshan
 * *************************************************/

package runner

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"time"
)

type Runner struct {
	interrupt chan os.Signal
	complete  chan error
	timeout   <-chan time.Time
	tasks     []func(int)
}

var errorTimeout = errors.New("received timeout. ")

var errorInterrupt = errors.New("receive interrupt. ")

func New(d time.Duration) *Runner {
	return &Runner{
		interrupt: make(chan os.Signal, 1),
		complete:  make(chan error),
		timeout:   time.After(d),
	}
}

func (runner *Runner) Add(tasks ...func(int)) {
	runner.tasks = append(runner.tasks, tasks...)
}

func (runner *Runner) Start() error {
	signal.Notify(runner.interrupt, os.Interrupt)

	go func() {
		runner.complete <- runner.run()
	}()

	select {
	case err := <-runner.complete:
		return err
	case <-runner.timeout:
		return errorTimeout
	}
}

func (runner *Runner) run() error {
	for id, task := range runner.tasks {
		if runner.goInterrupt() {
			return errorInterrupt
		}
		task(id)
	}
	return nil
}

func (runner *Runner) goInterrupt() bool {
	select {
	case <-runner.interrupt:
		signal.Stop(runner.interrupt)
		return true
	default:
		return false
	}
}

func TestRunner() {
	const timeout = 3 * time.Second
	log.Println("Starting work. ")

	runner := New(timeout)
	runner.Add(createTask(), createTask(), createTask())
	if err := runner.Start(); err != nil {
		switch err {
		case errorTimeout:
			log.Println("Terminating due to timeout. ")
			os.Exit(1)
		case errorInterrupt:
			log.Println("Terminating due to interrupt. ")
			os.Exit(2)
		}
	}

	log.Println("Process end. ")
}

func createTask() func(int) {
	return func(d int) {
		log.Printf("Processor - Task #%d. ", d)
		time.Sleep(time.Duration(d) * time.Second)
	}
}
