package hw05parallelexecution

import (
	"errors"
	"sync"
	"sync/atomic"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	ch := make(chan Task)
	wg := sync.WaitGroup{}

	var errCnt int32

	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			for task := range ch {
				if err := task(); err != nil {
					atomic.AddInt32(&errCnt, 1)
				}
			}
		}()
	}

	for _, task := range tasks {

	}
	// Place your code here.
	return nil
}
