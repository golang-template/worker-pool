package model

import (
	"fmt"
	"time"
)

// Task represents a unit of work for the worker pool.
type Task struct {
	ID      int
	Content string
}

// Worker represents a worker in the pool that processes tasks.
type Worker struct {
	ID         int
	TaskQueue  chan Task
	ResultChan chan string
}

// Start processes tasks from the task queue.
func (w Worker) Start() {
	for task := range w.TaskQueue {
		// Simulate work with time.Sleep
		time.Sleep(500 * time.Millisecond)

		// Simulate task completion or failure
		if task.ID%2 == 0 {
			w.ResultChan <- fmt.Sprintf("Worker %d successfully processed task %d: %s", w.ID, task.ID, task.Content)
		} else {
			w.ResultChan <- fmt.Sprintf("Worker %d failed to process task %d: %s", w.ID, task.ID, task.Content)
		}
	}
}
