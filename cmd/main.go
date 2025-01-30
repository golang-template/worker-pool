package main

import (
	"fmt"
	"time"
	"workerpool/model"
	"workerpool/worker"
)

func main() {
	tasks := CreateTasks(10)

	wp := worker.NewWorkerPool(3)

	go wp.Run(tasks)

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
		tasks = append(tasks, model.Task{ID: i, Content: "example Task" + fmt.Sprintf("%d", i)})
	}
	return tasks
}
