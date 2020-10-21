package kvm

import (
	"sync"
	"time"
)

var CronSch *CronScheduler

/*********************************************************************************
 * Configuration
 *********************************************************************************/
func ConfigCron() {
	c := NewCronScheduler(5)
	SetCronScheduler(c)
}

func (c *CronScheduler) Start(parentwg *sync.WaitGroup) {
	loop := 1
	c.Cr.Start()
	for {
		c.Run()
		time.Sleep(time.Duration(c.Interval * int(time.Second)))
		loop += 1
	}
	parentwg.Done()
}

