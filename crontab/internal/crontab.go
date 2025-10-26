package internal

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/robfig/cron/v3"
)

type Crontab struct {
	scheduler *cron.Cron
}

func NewCrontab() (*Crontab, error) {
	logger := log.New(os.Stderr, "crontab-", log.LUTC)

	scheduler := cron.New(
		cron.WithParser(
			cron.NewParser(
				cron.SecondOptional|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dow|cron.Descriptor,
			),
		),
		cron.WithLogger(cron.VerbosePrintfLogger(logger)),
	)

	_, err := scheduler.AddFunc("@every 5s", func() {
		fmt.Printf("This is your cron job. Current time is: %v\n", time.Now().Format(time.RFC3339Nano))
	})
	if err != nil {
		return nil, err
	}

	return &Crontab{
		scheduler: scheduler,
	}, nil
}

func (c *Crontab) Name() string {
	return "crontab"
}

func (c *Crontab) Mount(chi.Router) {
	c.scheduler.Start()
}

func (c *Crontab) Close() {
	c.scheduler.Stop()
}
