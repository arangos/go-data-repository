package config

import (
	"context"
	"github.com/robfig/cron/v3"
	"log"
	"time"
)

// Scheduler wraps a cron.Cron instance.
type Scheduler struct {
	cron *cron.Cron
}

// DocusignScheduleJob creates and starts the scheduler, scheduling your jobFunc
func DocusignScheduleJob(jobFunc func(context.Context)) (*cron.Cron, error) {
	// run once immediately
	go jobFunc(context.Background())
	c := cron.New(cron.WithChain(
		cron.Recover(cron.DefaultLogger),
	))

	// Wrap your jobFunc so we can pass a context (if you need one)
	_, err := c.AddFunc("@every 30m", func() {
		log.Printf("running scheduled job at %s", time.Now())
		jobFunc(context.Background())
	})
	if err != nil {
		return nil, err
	}

	c.Start()
	return c, nil
}

// Stop gracefully stops the scheduler, waiting for running jobs to finish.
func (s *Scheduler) Stop() context.Context {
	return s.cron.Stop()
}
