////////////////////////////////////////////////////////////////////////
//
// Copyright (C) 2021-2022 The HeyApple Authors.
//
// Use of this source code is governed by the GNU Affero General
// Public License as published by the Free Software Foundation,
// either version 3 of the License, or any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.
//
////////////////////////////////////////////////////////////////////////

package job

import "time"

// Job defines a task that the Scheduler can run at intervals.
type Job interface {
	Run()
}

// Scheduler handles automatic ticket creation and resolution.
type Scheduler struct {
	jobs     []Job
	tick     time.Duration
	delta    time.Duration
	shutdown chan bool
}

// NewScheduler returns a new Scheduler instance.
func NewScheduler(tick time.Duration, job ...Job) *Scheduler {
	return &Scheduler{
		jobs:     job,
		tick:     tick,
		delta:    tick,
		shutdown: make(chan bool),
	}
}

// Run starts the scheduler and listens for events.
func (s *Scheduler) Run() {
	for {
		select {
		case <-s.shutdown:
			s.shutdown <- true
			return
		case <-time.After(s.delta):
			break
		}

		start := time.Now()
		for _, j := range s.jobs {
			j.Run()
		}
		stop := time.Now()

		s.delta = s.tick - (stop.Sub(start))
	}
}

// Stop stops the Scheduler and shuts it down.
func (s *Scheduler) Stop() {
	s.shutdown <- true
	<-s.shutdown

	close(s.shutdown)
}
