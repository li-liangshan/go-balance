/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/27
 * @Last Modified by: liliangshan
 * *************************************************/

package cron

import "time"

type Cron interface {
	Next(t time.Time) time.Time
}

type Delay struct {
	Delay time.Duration
}

func Every(duration time.Duration) Delay {
	if duration < time.Second {
		duration = time.Second
	}

	return Delay{
		Delay: duration - time.Duration(duration.Nanoseconds())%time.Second,
	}
}

func (delay Delay) Next(t time.Time) time.Time {
	return t.Add(delay.Delay - time.Duration(t.Nanosecond())*time.Nanosecond)
}
