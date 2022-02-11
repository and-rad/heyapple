package mock

import "time"

type Job struct {
	Done  bool
	Delay time.Duration
}

func NewJob(delay time.Duration) *Job {
	return &Job{Delay: delay}
}

func (j *Job) Run() {
	time.Sleep(j.Delay)
	j.Done = true
}
