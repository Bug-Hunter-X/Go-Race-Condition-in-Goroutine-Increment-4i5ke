# Go Race Condition Example

This repository demonstrates a common race condition in Go.  The code involves multiple goroutines incrementing a shared variable without proper synchronization mechanisms (mutexes or atomic operations). This can result in incorrect and non-deterministic results.

## Bug Description
The `bug.go` file contains a program that launches multiple goroutines, each incrementing a shared integer variable.  Due to the lack of synchronization, the final value of the variable is not guaranteed to be 10. The output will vary on each run.

## Solution
The `bugSolution.go` file demonstrates how to fix the race condition using a mutex. This ensures that only one goroutine can access and modify the shared variable at a time, producing consistent and predictable results.