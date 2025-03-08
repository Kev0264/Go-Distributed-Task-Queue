# Go-Distributed-Task-Queue

This is a set of packages that demonstrates concurrency, networking, and system design. The packages consist of multiple workers that process tasks from a central queue, demonstrating key concepts like goroutines, channels, networking, and fault tolerance

## Task Producer

The producer generates tasks and sends them to the queue

## Task Queue

The queue holds tasks and distributes them to workers

## Worker

Workers pull tasks from the queue and process them

# Running the system

1. Start the Task Queue:

```
cd queue
go run queue.go
```

2. Start the Workers:

```
cd worker
go run worker.go
```

3. Start the Task Producer:

```
cd producer
go run producer.go
```
