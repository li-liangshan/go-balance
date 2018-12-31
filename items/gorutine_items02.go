/* **************************************************
 * @Author: liliangshan
 * @Date: 2018/12/31
 * @Last Modified by: liliangshan
 * *************************************************/

package items

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counter int64
	wg sync.WaitGroup
)

func Items02() {
	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()

	fmt.Println("Final counter:", counter)
}

func incCounter(id int) {
	defer wg.Done()

	for count := 0; count < 3; count++ {
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}
