/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/27
 * @Last Modified by: liliangshan
 * *************************************************/

package times

import (
	"fmt"
	"time"
)

func TestTimes() {
	go func() {
		var timer *time.Timer
		for   {
			select {
			case <- func() <-chan time.Time {
				if timer==nil {
					timer=time.NewTimer(time.Millisecond)
				}else {
					timer.Reset(time.Millisecond)
				}
				return timer.C
			}():
				fmt.Println("time out")
				break
			}
		}
	}()

}
