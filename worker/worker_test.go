/* **************************************************
 * @Author: liliangshan
 * @Date: 2019/1/1
 * @Last Modified by: liliangshan
 * *************************************************/

package worker

import (
	"log"
	"sync"
	"testing"
	"time"
)

var names = []string{
	"steve", "bob", "mary", "therese", "jason",
}

type namePrinter struct {
	name string
}

func (printer *namePrinter) Task() {
	log.Println(printer.name)
	time.Sleep(time.Second)
}

func TestWorker(t *testing.T) {
	pool := New(2)
	var wg sync.WaitGroup
	wg.Add(100 * len(names))

	for i := 0; i < 100; i++ {
		for _, name := range names {
			np := namePrinter{
				name: name,
			}

			go func() {
				pool.Run(&np)
				wg.Done()
			}()
		}
	}

	wg.Wait()
	pool.Shutdown()
}
