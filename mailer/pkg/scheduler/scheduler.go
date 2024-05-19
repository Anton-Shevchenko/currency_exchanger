package scheduler

import (
	"github.com/go-co-op/gocron"
	"time"
)

type JobInterface interface {
	Run()
}

type Scheduler struct {
	scheduler *gocron.Scheduler
	job       JobInterface
}

func NewScheduler(job JobInterface) *Scheduler {
	s := gocron.NewScheduler(time.UTC)
	return &Scheduler{
		scheduler: s,
		job:       job,
	}
}

func (s *Scheduler) Daily() {
	s.scheduler.Every(1).Day().At("00:00").Do(s.job.Run)
	s.scheduler.StartAsync()
}
