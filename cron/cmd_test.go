/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/27
 * @Last Modified by: liliangshan
 * *************************************************/

package cron

import (
	"fmt"
	"testing"
	"time"
)

func TestCmd(t *testing.T) {
	tim := time.Now()
	res := 1 << uint(tim.Month()) & uint(1)
	fmt.Println("res:", res)
	t.Log("res:", res)
	t.Log("month:", uint(tim.Month()))
}
