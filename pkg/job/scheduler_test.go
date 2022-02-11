package job

import (
	"heyapple/internal/mock"
	"testing"
	"time"
)

const (
	deltaThreshold = 5
)

func TestScheduler_Run(t *testing.T) {
	for idx, data := range []struct {
		jobs  []Job
		tick  time.Duration
		delta time.Duration
	}{
		{ //00// no jobs, stop immediately
			delta: 0,
		},
		{ //00//
			jobs:  []Job{mock.NewJob(0)},
			delta: 0,
		},
		{ //00// tick not set, delta as long as all jobs take
			jobs: []Job{
				mock.NewJob(time.Millisecond * 200),
				mock.NewJob(time.Millisecond * 100),
			},
			delta: time.Millisecond * 300,
		},
		{ //00//
			jobs: []Job{
				mock.NewJob(0),
				mock.NewJob(0),
			},
			tick:  time.Millisecond * 300,
			delta: time.Millisecond * 300,
		},
		{ //00//
			jobs: []Job{
				mock.NewJob(0),
				mock.NewJob(time.Millisecond * 150),
			},
			tick:  time.Millisecond * 300,
			delta: time.Millisecond * 150,
		},
	} {
		s := NewScheduler(data.tick, data.jobs...)
		go s.Run()

		time.Sleep(data.tick + time.Millisecond*200)
		s.Stop()

		if (s.delta - data.delta) > deltaThreshold {
			t.Errorf("test case %d: delta mismatch \nhave: %v \nwant: %v", idx, s.delta, data.delta)
		}

		for i, j := range data.jobs {
			if x, ok := j.(*mock.Job); ok && !x.Done {
				t.Errorf("test case %d-%d: job not done: %v", idx, i, x)
			}
		}
	}
}
