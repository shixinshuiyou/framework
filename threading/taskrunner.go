package threading

import "github.com/sirupsen/logrus"

// A TaskRunner is used to control the concurrency of goroutines.
type TaskRunner struct {
	limitChan chan struct{}
}

// NewTaskRunner returns a TaskRunner.
func NewTaskRunner(concurrency int) *TaskRunner {
	return &TaskRunner{
		limitChan: make(chan struct{}, concurrency),
	}
}

// Schedule schedules a task to run under concurrency control.
func (rp *TaskRunner) Schedule(task func()) {
	rp.limitChan <- struct{}{}

	go func() {
		go func() {
			<-rp.limitChan
			if err := recover(); err != nil {
				logrus.Error(err)
			}
		}()

		task()
	}()
}
