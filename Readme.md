# Concurrency Patterns in Go lang

## Generator Pattern
   - The Generator Pattern in go is used to create a stream of values that can be consumed lazily. The main use cases for thic
     pattern is in streaming Data processing, Lazy loading of Large data, polling APIs
     [implementation](./generator/generator.go)

   - Key Benefits
        - ✅ Lazy Evaluation: Values are produced only when needed.
        - ✅ Concurrency Friendly: Goroutines allow background computation.
        - ✅ Efficient Memory Usage: Only one value is held in memory at a time.

## Fan-out/Fan-In
   - Fan-In and Fan-out patterns in go help manage workloads effiiciently using goroutines and channels.
        - Fan-out Pattern (Distributing the work)
            The Fan-out pattern is used to distribute the work across multiple goroutines to improve performance.
            Distribute the work among multiple workers.

            ✅ Improves throughput by using multiple workers.

        - Fan-in Patterns (Merging multiple channels)
            The Fan-in pattern combines multiple input channels into single output channel.
            Used to aggregate results from multiple sources into single stream.
           - Key Benefits
             - ✅ Simplifies concurrent data processing.
             - ✅ Combines multiple sources into one stream.

## Worker pool pattern
   - The worker pool pattern is used to limit the number of worker goroutines processing the tasks at the same time. This
      prevents resources usage and prevents excessive usage of goroutines.

      - How it works
        - A pool of goroutines listnes for tasks from a shared job queue.
        - Each worker process the tasks concurrently.
        
      - When to use worker pool pattern
        - Batch processing.
        - API request Handling (rate limiting concurrent requests).

      - Key Notes:
        - Senders close channels, NOT receivers. The main goroutine (which sends jobs) should close the jobs channel.
        - Workers (receivers) just read from the channel and stop when it's closed.

      - Key Benefits:
         - ✅ Efficient resource usage (controls goroutine count).
         - ✅ Parallel processing improves performance.
         - ✅ Prevents excessive goroutines from being created.

## Pipeline Pattern
   - The pipeline pattern in Go is used to process data in multiple stages using goroutines and channels. 
       This improves efficiency by enabling concurrent processing while keeping the code readable and scalable.

        - How it works
            - A sequence of trnsformations, each stage runs it's own goroutine and receives inout from the previous stage.
            - Channels connect and pass the data from one stage to the next.
            - The pipeline process the data concurrently, improving efficiency.
        
        - When to use pipeline pattern
            - Stream processing.
            - Batch data transformation.
            - Parallel Processing.

        - Key Benefits:
           - ✅ Concurrent execution without blocking.
           - ✅ Scalable—can add more stages easily.
           - ✅ Memory efficient—no need to store all results in memory.

## Pub-Sub Pattern
   - The Pub-sub (Publish - subscribe) pattern in a messaging pattern where publishers send messsages to multiple sbscribers
      i.e broad casts messages to multiple subscribers. pub-sub communication happens with the help of buffered channel, where subscribers listens to the channels for any updates.

        - How it works
          - Publisher send messages to channel.
          - Subscribers listen to that channel and process messages.
          - Multiple subsribers can recieve the same message simultaneously.
        
        - Key Benefits
          - ✅ Decouples publishers and subscribers (flexible architecture).
          - ✅ Scalable (multiple subscribers can listen to the same topic).
          - ✅ Supports dynamic subscription/unsubscription (clients can join/leave at any time).
        
        - Use cases;
          - Event-driven systems.
          - Message queues.

## Semaphore pattern
   - A Semaphore pattern is used to limit the number of concurrent operations in a system. It helps prevent resource   exhaustion by allowing only fixed number of goroutines to proceed simultaneously.
        
        - How it works
          - A semaphore maintains a counter representing available resources.
          - Goroutines aquire a permit before proceeding.
          - if no permit are available, goroutines must wait until one is released.
          - when done, goroutines release the permit, allowing other to proceed.
        
        🔹 Binary Semaphore (count = 1) → Works like a sync.Mutex (only one goroutine at a time).
        🔹 Counting Semaphore (count > 1) → Allows multiple goroutines to run concurrently.

        - Use cases:
          - Rate limiting.
          - Controlling access to limited number of resources.

        - Both semaphore and worker pool patterns help control concurrency, but they serve different purposes and work differently.
        - Any goroutine can run when permit is available.
        - Does not create a set number of workers.
        - In worker pool pattern a fixed number of goroutines are created that process tasks from a queue.

        ✅ Semaphore Pattern → Controls how many goroutines run at once, but they work independently.
        ✅ Worker Pool Pattern → Pre-creates a fixed number of worker goroutines that process jobs from a queue.

        - Rate limiting API calls (e.g., only 5 API calls at a time)	Semaphore.
        - Controlling concurrent access to a database or resource	Semaphore.
        - Processing multiple tasks efficiently with a fixed number of workers	Worker Pool.
        - Task queue processing (e.g., background jobs, file processing, batch processing)	Worker Pool.

## Heartbeat pattern
    - The Heartbeat Pattern is used in distributed systems and concurrent applications to monitor the health of a service or process. It involves periodic signals ("heartbeats") sent from one component (e.g., a worker, service, or system) to another (e.g., a monitor, orchestrator, or supervisor). If no heartbeat is received within a given time, the system assumes failure and takes action.

      - How it works
        - A worker/service sends heartbeats periodically.
        - A monitor listens for heartbeats.
        - If the monitor stops receiving heartbeats within a timeout period, it assumes failure and reacts accordingly.

      

