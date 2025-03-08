package main

import (
	"encoding/json"
	"fmt"
	"go-distributed-task-queue/shared"
	"net/http"
	"sync"
)

var (
	taskQueue []shared.Task
	queueLock sync.Mutex
)

func encueueHandler(w http.ResponseWriter, r *http.Request) {
	var task shared.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		http.Error(w, "Invalid task", http.StatusBadRequest)
		return
	}

	queueLock.Lock()
	taskQueue = append(taskQueue, task)
	queueLock.Unlock()

	fmt.Printf("Enqueued task %d\n", task.ID)
	w.WriteHeader(http.StatusOK)
}

func dequeueHandler(w http.ResponseWriter, r *http.Request) {
	queueLock.Lock()
	defer queueLock.Unlock()

	if len(taskQueue) == 0 {
		http.Error(w, "No task available", http.StatusNotFound)
		return
	}

	task := taskQueue[0]
	taskQueue = taskQueue[1:]

	json.NewEncoder(w).Encode(task)
	fmt.Fprintf(w, "Task dequeued: %d\n", task.ID)
}

func main() {
	http.HandleFunc("/enqueue", encueueHandler)
	http.HandleFunc("/dequeue", dequeueHandler)

	fmt.Println("Queue service started at :8080")
	http.ListenAndServe(":8080", nil)
}
