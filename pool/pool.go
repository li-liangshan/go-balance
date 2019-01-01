/* **************************************************
 * @Author: liliangshan
 * @Date: 2019/1/1
 * @Last Modified by: liliangshan
 * *************************************************/

package pool

import (
	"errors"
	"io"
	"log"
	"sync"
)

type Pool struct {
	mutex     sync.Mutex
	resources chan io.Closer
	factory   func() (io.Closer, error)
	closed    bool
}

var ErrorPoolClosed = errors.New("Pool has been closed. ")

func New(fn func() (io.Closer, error), size uint) (*Pool, error) {
	if size <= 0 {
		return nil, errors.New("size value too small. ")
	}

	return &Pool{
		factory:   fn,
		resources: make(chan io.Closer, size),
	}, nil
}

func (pool *Pool) Acquire() (io.Closer, error) {
	select {
	case r, ok := <-pool.resources:
		log.Println("Acquire: ", "Shared resource")
		if !ok {
			return nil, ErrorPoolClosed
		}
		return r, nil
	default:
		log.Println("Acquire: ", "New resource")
		return pool.factory()
	}
}

func (pool *Pool) Release(r io.Closer) {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()

	if pool.closed {
		r.Close()
		return
	}

	select {
	case pool.resources <- r:
		log.Println("Release: ", "In Queue")
	default:
		log.Println("Release: ", "Closing")
		r.Close()
	}
}

func (pool *Pool) Close() {
	pool.mutex.Lock()
	defer pool.mutex.Unlock()

	if pool.closed {
		return
	}

	pool.closed = true
	close(pool.resources)
	for r := range pool.resources {
		r.Close()
	}
}
