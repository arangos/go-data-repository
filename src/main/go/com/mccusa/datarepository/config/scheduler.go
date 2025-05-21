package config

import (
	"github.com/go-co-op/gocron"
	"time"
)

// NewTaskScheduler creates and starts a scheduler with a worker pool size of 5.
func NewTaskScheduler() *gocron.Scheduler {
	scheduler := gocron.NewScheduler(time.UTC)
	// Limit to 5 concurrent jobs
	scheduler.SetMaxConcurrentJobs(5, gocron.RescheduleMode)
	scheduler.StartAsync()
	return scheduler
}
