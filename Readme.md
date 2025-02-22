# Concurrency Patterns in Go lang

## Generator Pattern
   - The Generator Pattern in go is used to create a stream of values that can be consumed lazily. The main use cases for thic
     pattern is in streaming Data processing, Lazy loading of Large data, polling APIs
     [implementation](./generator/generator.go)

   - Key Benefits
    ✅ Lazy Evaluation: Values are produced only when needed.
    ✅ Concurrency Friendly: Goroutines allow background computation.
    ✅ Efficient Memory Usage: Only one value is held in memory at a time.

## Fan-out/Fan-In
   - Fan-In and Fan-out patterns in go help manage workloads effiiciently using goroutines and channels.
        - Fan-out Pattern (Distributing the work)
            The Fan-out pattern is used to distribute the work across multiple goroutines to improve performance.
            Distribute the work among multiple workers.

            ✅ Improves throughput by using multiple workers.

        - Fan-in Patterns (Merging multiple channels)
            The Fan-in pattern combines multiple input channels into single output channel.
            Used to aggregate results from multiple sources into single stream.

            ✅ Simplifies concurrent data processing.
            ✅ Combines multiple sources into one stream.

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
      
      ✅ Efficient resource usage (controls goroutine count).
      ✅ Parallel processing improves performance.
      ✅ Prevents excessive goroutines from being created.

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

        ✅ Concurrent execution without blocking.
        ✅ Scalable—can add more stages easily.
        ✅ Memory efficient—no need to store all results in memory.
