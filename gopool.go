package gopool

import (
	"sync"
)

// workerPool is a collection of goroutines in a way the number of concurrent goroutines processing requests
// does not exceed the specified maximum.
type workerPool struct {
	maxWorker         uint
	waitgroup         sync.WaitGroup
	queuedTaskChannel chan Job
}

// Job is an interface that specifies the function for a worker to execute.
type Job interface {
	Do()
}

// WorkerPool is an interface that specifies the functions to handle the pool of worker goroutines.
type WorkerPool interface {
	AddTask(task Job)
	Shutdown()
}

// NewWorkerPool is a factory function to create and start a pool of worker goroutines.
//
// The maxWorker parameter specifies the maximum number of workers that can execute tasks concurrently.
// The capacity parameter specifies the maximum number of tasks that can be put in a queue.
func NewWorkerPool(maxWorker, capacity uint) WorkerPool {
	wp := &workerPool{
		maxWorker:         maxWorker,
		waitgroup:         sync.WaitGroup{},
		queuedTaskChannel: make(chan Job, capacity),
	}
	wp.start()
	return wp
}

// AddTask enqueues a Job for a worker to execute.
//
// It will not block regardless of the number of tasks. Each task is immediately given to an available or a new worker.
// If there are no available workers and the maximum number of workers is already created, the task is put in a waiting queue.
func (wp *workerPool) AddTask(task Job) {
	wp.waitgroup.Add(1)
	wp.queuedTaskChannel <- task
}

// start executes the queued tasks by an available or a new worker goroutines concurrently.
func (wp *workerPool) start() {
	for i := uint(0); i < wp.maxWorker; i++ {
		go func() {
			for task := range wp.queuedTaskChannel {
				task.Do()
				wp.waitgroup.Done()
			}
		}()
	}
}

// Shutdown closes task queue and waits for currently running tasks to finish.
func (wp *workerPool) Shutdown() {
	wp.waitgroup.Wait()
	close(wp.queuedTaskChannel)
}
