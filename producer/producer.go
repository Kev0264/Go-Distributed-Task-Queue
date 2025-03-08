package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-distributed-task-queue/shared"
	"net/http"
	"time"
)

func main() {
	for i := 1; i < 10; i++ {
		task := shared.Task{ID: i, Data: fmt.Sprintf("Task data %d", i)}
		jsonData, _ := json.Marshal(task)

		// Send task to the queue
		resp, err := http.Post("http://localhost:8080/enqueue", "application/json", bytes.NewBuffer(jsonData))
		if err != nil {
			fmt.Println("Error sending task to the queue: ", err)
			continue
		}
		defer resp.Body.Close()

		fmt.Println("Task sent to the queue: ", task.ID)
		time.Sleep(1 * time.Second) // Simulate task generation delay
	}
}
