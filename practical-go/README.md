# Practical Concurrency in Go Notes
## General Notes
- - How do you handle y processes with x cores when y > x
	- For example I have  `ps aux | wc -l` will show you how many processes you have, and `lscpu` will show you how many CPUs you have
- Concurrency is cooking a meal
	- You preheat the oven and while it's preheating you wash and slice up vegetables
	- While the vegetables are cooking, you can clean up and get ready to plate up
	- Context switching between tasks
- Parallelism is multiple cooks in the kitchen.
	- Concurrency enables the potential for parallelism because is multiple cooks can divy up the work already into non-conflicting parts, they can work at the same time without conflict or collision
- Latency is a bugbear of a problem
	- ![](https://pbs.twimg.com/media/D2IYawBXcAAhGje?format=png&name=900x900)
- The Go scheduler knows when a process is going for a long journey on the net, so it can schedules based on the wait time

## Goroutine Code Examples
- This example shows a bug you want to avoid along with the two ways to fix it: [Example](https://go.dev/play/p/WQMYZ-7xifM)
- This example shows how to do this with channels: [Example](https://go.dev/play/p/BHVhBD17cJA)
- This examples shows using channels for synchronization: [Example](https://go.dev/play/p/Nz6vX_x3Sg9)
- To run taxi_check:
  - `cd taxi_check/` and run `go run taxi_check_sequential.go` and then `go run taxi_check.go` to see the difference concurrency makes.
  - Look at the code inside `taxi_check.go` for notes on the semaphore pattern to control how many goroutines we have running at any given time