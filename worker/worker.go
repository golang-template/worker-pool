package worker

import (
	"sync"
	"workerpool/model"
)

// WorkerPool manages a set of workers and coordinates task execution.
type WorkerPool struct {
	NumWorkers int
	TaskQueue  chan model.Task
	ResultChan chan string
}

// NewWorkerPool creates a new worker pool instance.
func NewWorkerPool(numWorkers int) *WorkerPool {
	return &WorkerPool{
		NumWorkers: numWorkers,
		TaskQueue:  make(chan model.Task, numWorkers),
		ResultChan: make(chan string, numWorkers),
	}
}

// Run starts the worker pool, dispatches tasks to workers, and collects results.
func (wp *WorkerPool) Run(tasks []model.Task) {
	var wg sync.WaitGroup

	// Start workers
	for i := 1; i <= wp.NumWorkers; i++ {
		worker := NewWorker(i, wp.TaskQueue, wp.ResultChan)
		wg.Add(1)
		go func(w model.Worker) {
			defer wg.Done()
			w.Start() // Call Start on model.Worker
		}(worker)
	}

	// Enqueue tasks
	for _, task := range tasks {
		wp.TaskQueue <- task
	}

	// Close task queue to signal no more tasks
	close(wp.TaskQueue)

	// Wait for all workers to finish
	wg.Wait()

	// Close result channel to prevent blocking on results
	close(wp.ResultChan)
}

// NewWorker creates a new worker instance.
func NewWorker(id int, taskQueue chan model.Task, resultChan chan string) model.Worker {
	return model.Worker{
		ID:         id,
		TaskQueue:  taskQueue,
		ResultChan: resultChan,
	}
}
