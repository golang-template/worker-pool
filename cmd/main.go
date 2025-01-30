package main

import (
	"fmt"
	"time"
	"workerpool/model"
	"workerpool/worker"
)

func main() {
	// Create tasks
	tasks := CreateTasks(10)

	// Create a worker pool with 3 workers
	wp := worker.NewWorkerPool(3)

	// Run the worker pool
	go wp.Run(tasks)

	// Collect and print results from workers
	for result := range wp.ResultChan {
		fmt.Println(result)
	}

	// Add a delay to ensure workers finish before the main function exits
	time.Sleep(3 * time.Second)
}

// CreateTasks generates a list of tasks for the worker pool to process.
func CreateTasks(n int) []model.Task {
	var tasks []model.Task
	for i := 1; i <= n; i++ {
		tasks = append(tasks, model.Task{ID: i, Content: "Task #" + fmt.Sprintf("%d", i)})
	}
	return tasks
}
