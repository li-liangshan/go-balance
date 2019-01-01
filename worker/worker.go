/* **************************************************
 * @Author: liliangshan
 * @Date: 2019/1/1
 * @Last Modified by: liliangshan
 * *************************************************/

package worker

import "sync"

type Worker interface {
	Task()
}

type Pool struct {
	worker chan Worker
	wg     sync.WaitGroup
}

func New(maxGoroutines int) *Pool {
	pool := Pool{
		worker: make(chan Worker),
	}

	pool.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for worker := range pool.worker {
				worker.Task()
			}
			pool.wg.Done()
		}()
	}

	return &pool
}

func (pool *Pool) Run(worker Worker) {
	pool.worker <- worker
}

func (pool *Pool) Shutdown() {
	close(pool.worker)
	pool.wg.Wait()
}
