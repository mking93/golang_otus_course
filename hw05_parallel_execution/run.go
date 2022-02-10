package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	var (
		errCount int

		mu sync.Mutex
		wg sync.WaitGroup
	)

	wg.Add(n)
	tasksCh := make(chan Task, len(tasks))

	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()

			isDone := false
			for task := range tasksCh {
				err := task()

				mu.Lock()
				if errCount >= m {
					isDone = true
				}
				if err != nil {
					errCount++
				}
				mu.Unlock()

				if isDone {
					return
				}
			}
		}()
	}

	for _, task := range tasks {
		tasksCh <- task
	}
	close(tasksCh)

	wg.Wait()
	if errCount >= m {
		return ErrErrorsLimitExceeded
	}

	return nil
}
