package main

import "fmt"

// Реализовать паттерн «адаптер» на любом примере.

type CommonJob interface {
	Do()
}

type User struct {
}

func (u *User) DoCommonJob(j CommonJob) {
	j.Do()
}

type Job1 struct {
}

func (j *Job1) Do() {
	fmt.Println("Do job 1")
}

type HardTask struct{}

func (h *HardTask) WorkHard() {
	fmt.Println("Working hard")
}

type HardTaskAdapter struct {
	*HardTask
}

func (h *HardTaskAdapter) Do() {
	h.WorkHard()
}

func main() {
	user := &User{}
	job1 := &Job1{}
	user.DoCommonJob(job1)
	hardTask := &HardTask{}
	hardTaskAdapter := &HardTaskAdapter{hardTask}
	user.DoCommonJob(hardTaskAdapter)
}
