package main

import (
	"TaskQueue/asyncq"
	"log"
)

type Test struct {
	Index int
}

func NewTest(i int) *Test {
	return &Test{Index: i}
}

func (s *Test) Perform() {
	log.Printf("Perform index = %d\n", s.Index)
}

func main() {
	//log.Println("Program is started")

	asyncq.StartTaskDispatcher(5)

	for i := 0; i < 108; i++ {
		test := NewTest(i)
		asyncq.TaskQueue <- test
	}

	var count = 0
	for len(asyncq.TaskQueue) > 0 {
		count++
	}

	log.Println("count of counter = ", count)
	log.Println("count of tasks in TaskQueue = ", len(asyncq.TaskQueue))
	//time.Sleep(5 * time.Second)
}
