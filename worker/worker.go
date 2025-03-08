package main

import (
	"encoding/json"
	"fmt"
	"go-distributed-task-queue/shared"
	"net/http"
	"time"
)

func worker(workerID int) {
	for {
		resp, err := http.Get("http://localhost:8080/dequeue")
		if err != nil {
			fmt.Printf("Worker %d: Error fetching task: %v\n", workerID, err)
			time.Sleep(2 * time.Second) // Retry after delay
			continue
		}

		if resp.StatusCode == http.StatusNotFound {
			fmt.Printf("Worker %d: No task available\n", workerID)
			time.Sleep(2 * time.Second) // Wait for new tasks
			continue
		}

		var task shared.Task
		if err := json.NewDecoder(resp.Body).Decode(&task); err != nil {
			fmt.Printf("Worker %d: Error decoding task: %v\n", workerID, err)
			resp.Body.Close()
			continue
		}
		resp.Body.Close()

		fmt.Printf("Worker %d processing task %d: %s\n", workerID, task.ID, task.Data)
		time.Sleep(3 * time.Second) // Simulate task processing
	}
}

func main() {
	for i := 1; i <= 3; i++ {
		go worker(i)
	}

	// Keep the main goroutine alive
	select {}
}