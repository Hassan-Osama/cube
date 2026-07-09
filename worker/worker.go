package worker

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"

	"cube/task"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]*task.Task
	TaskCount int
}

func (w *Worker) CollectStats() {
	fmt.Println("i am collecting stats")
}

func (w *Worker) RunTask() {
	fmt.Println("i am running task")
}

func (w *Worker) StartTask() {
	fmt.Println("i am running task")
}

func (w *Worker) StopTask() {
	fmt.Println("iam stopping task")
}
