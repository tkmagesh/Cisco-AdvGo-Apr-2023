package main

import (
	"fmt"
	"os"
	"time"

	"runner-demo/runner"
)

func main() {
	/*
		initialize the runner with a timeout
		add multiple tasks to the runner
		Start the runner
		The tasks are executed sequentially
		if all the tasks are completed within the given time, report "Success"
		if the tasks are not completed within the given time, report "timeout"
		exit if the execution is interrupted by an OS interrupt
	*/
	fmt.Printf("Process %d started....\n", os.Getpid())

	timeout := 15 * time.Second
	r := runner.New(timeout)
	r.Add(createTask(8))
	r.Add(createTask(6))
	r.Add(createTask(5))

	if err := r.Start(); err != nil {
		switch err {
		case runner.ErrTimeout:
			fmt.Println("tasks timed out")
		default:
			fmt.Println("unknown error : ", err)
		}
	} else {
		fmt.Println("success")
	}

}

func createTask(t int) func(int) {
	return func(id int) {
		fmt.Printf("Processing task # %d\n", id)
		time.Sleep(time.Duration(t) * time.Second)
	}
}
