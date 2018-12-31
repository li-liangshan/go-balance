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
)

var (
	counter1 int64
	wg1 sync.WaitGroup
	mutex sync.Mutex
)

func Items03() {
	wg1.Add(2)

	go incCounter03(1)
	go incCounter03(2)

	wg1.Wait()

	fmt.Println("Final counter:", counter1)
}

func incCounter03(id int) {
	defer wg1.Done()

	for count := 0; count < 3; count++ {
		mutex.Lock()
		{
			value := counter1
			runtime.Gosched()
			value++
			counter1 = value
		}
		mutex.Unlock()
	}
}
