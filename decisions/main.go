package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func textFromInput(reader *bufio.Reader) string {
	task, _ := reader.ReadString('\n')
	return strings.Replace(task, "\n", "", -1)
}

func main() {
	fmt.Println("Decisions decisions. What all do you have to decide between? ")
	reader := bufio.NewReader(os.Stdin)
	task := textFromInput(reader)
	var tasks []string
	tasks = append(tasks, task)
	for {
		fmt.Println("anything else? (enter 'done' or 'DONE' to finish) ")
		task = textFromInput(reader)
		if strings.ToUpper(task) == "DONE" {
			break
		}
		tasks = append(tasks, task)
	}
	for len(tasks) > 0 {
		fmt.Printf("can you do '%s' (Y/N)? ", tasks[0])
		response := textFromInput(reader)
		if strings.ToUpper(response) == "Y" {
			fmt.Println("Well go do that! You've got this! Cha La Head Cha La!")
			os.Exit(0)
		} else {
			copy(tasks[0:], tasks[1:])   // Shift a[i+1:] left one index.
			tasks[len(tasks)-1] = ""     // Erase last element (write zero value).
			tasks = tasks[:len(tasks)-1] // Truncate slice.
		}
	}
	fmt.Println("Looks like you need to find a task you can do right now. Try again?")
}
