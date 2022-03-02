package main

import (
	"fmt"
	"time"
)

type Job struct {
	ID int

	State string

	done chan bool
}

func (j *Job) Execute() {
	fmt.Println("executing:\t", j.ID)
	j.State = "running"
	fmt.Println(j.State)
	time.Sleep(100 * time.Millisecond)
	j.State = "done"
	close(j.done)
}

func (j Job) Wait() {
	<-j.done
}

//consumer
func worker(ch chan *Job) {
	for job := range ch {
		job.Execute()
	}
}

func main() {
	num := 3
	ch := make(chan *Job)
	for i := 0; i < num; i++ {
		go worker(ch)
	}

	//Producer
	for i := 0; i < num; i++ {
		j := &Job{
			ID:    i,
			State: "ready",
			done:  make(chan bool),
		}
		ch <- j
		j.Wait()
		fmt.Println("Finished; final state is:\t", j.State)
	}

}
