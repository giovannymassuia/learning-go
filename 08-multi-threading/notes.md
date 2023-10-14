## Go Multi Threading


### Processes

- Processes in go are called goroutines.
- Goroutines are lightweight threads managed by the go runtime.
- Processes in go are independent of each other.

### Threads

- Threads are the smallest unit of execution in a process.
- Threads in the same process share the same memory space.

### Concurrency and Mutex

- Concurrency is the ability to run multiple processes at the same time.
  - Concurrency is not parallelism.
  - Concurrency is about dealing with lots of things at once.
  - Parallelism is about doing lots of things at once.

- Mutex is a mutual exclusion lock.
  - It is used to synchronize access to a shared resource.

### Concurrency vs Parallelism vs Go Routines

- Concurrency is about dealing with lots of things at once.
  - Real life examples: A person talking on the phone while typing on the computer.
- Parallelism is about doing lots of things at once.
  - Real life example: A person talking on the phone and another person typing on the computer.
- Go routines are lightweight threads managed by the go runtime.
  - Go routines are:
    - Concurrent
  - Go routines are not:
    - Parallel
    - Threads
    - Processes
  - When running with multiple cores, go routines can run in parallel.
    - This is not guaranteed.
    
### Go Routines

- Go routines are lightweight threads managed by the go runtime.
- Go routines are concurrent.
- Go routines are not parallel.
- Go routines are not threads.

### Channels

- Channels are used to communicate between go routines.
  - Go routines in different cores can communicate with each other using channels.

### Cost of Threads

- Threads are expensive.
- Threads have a large memory footprint.
- Context switching between threads is expensive.
  - Context switching is the process of storing and restoring the state of a thread.
- Syscalls are expensive.
  - Syscalls are calls to the OS.

- Asking for a new thread to the OS step by step:
  - Program asks for a new thread. (as a process)
  - OS allocates memory for the thread.
  - OS sets up the thread.

### Schedulers

- Schedulers are responsible for scheduling threads on the CPU.
- Schedulers are part of the OS.

- Types of schedulers:
  - Preemptive
    - The scheduler can interrupt a thread at any time.
    - The scheduler can interrupt a thread while it is running.
    - The scheduler can interrupt a thread while it is waiting.
  - Cooperative
    - The scheduler can only interrupt a thread while it is waiting.
    - The scheduler cannot interrupt a thread while it is running.
  - Hybrid
    - The scheduler can interrupt a thread while it is running.
    - The scheduler cannot interrupt a thread while it is waiting.

### Go Scheduler

- The go scheduler is a preemptive scheduler.
- The go scheduler can interrupt a thread at any time.
- The go scheduler can interrupt a thread while it is running.
- The go scheduler can interrupt a thread while it is waiting.
- The go scheduler is part of the go runtime.
- The go scheduler is not part of the OS.
- The go scheduler is not a thread, not a process, not a go routine.

### Go and its green threads

- Go uses green threads.
- Green threads are managed by the go runtime.

- How does it work?
  - The go runtime creates a number of OS threads.
  - The go runtime schedules go routines on the OS threads.
  - The go runtime schedules go routines on the OS threads using the `go scheduler`.
  - By default, go schedules is cooperative, but it can be preemptive.
  - The go scheduler can interrupt a go routine at any time.
  - The go scheduler can interrupt a go routine while it is running.
  - The go scheduler can interrupt a go routine while it is waiting.

- Cost of OS threads vs Go routines
  - OS threads: ~1mb
  - Go routines: ~2kb