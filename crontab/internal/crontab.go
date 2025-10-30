package internal

import (
	"log"
	"os"
	"time"

	"github.com/robfig/cron/v3"
	"github.com/titpetric/platform"
)

type Crontab struct {
	platform.UnimplementedModule

	scheduler *cron.Cron
}

func NewCrontab() *Crontab {
	logger := log.New(os.Stderr, "crontab-", log.LUTC)

	scheduler := cron.New(
		cron.WithParser(
			cron.NewParser(
				cron.SecondOptional|cron.Minute|cron.Hour|cron.Dom|cron.Month|cron.Dow|cron.Descriptor,
			),
		),
		cron.WithLogger(cron.VerbosePrintfLogger(logger)),
	)

	return &Crontab{
		scheduler: scheduler,
	}
}

func (c *Crontab) Name() string {
	return "crontab"
}

func (c *Crontab) Start() error {
	_, err := c.scheduler.AddFunc("@every 5s", func() {
		log.Printf("This is your cron job starting.")
		time.Sleep(3 * time.Second)
		log.Printf("Cron job exiting after 3 secs.")
	})
	if err != nil {
		return err
	}

	c.scheduler.Start()
	return nil
}

func (c *Crontab) Stop() error {
	<-c.scheduler.Stop().Done()
	return nil
}
