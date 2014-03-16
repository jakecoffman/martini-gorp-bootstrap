package main

import "sync"

type Task struct {
	Id     int    `json:"id"`
	Script string `json:"script"`
}

type TaskService interface {
	Get(id int) *Task
	GetAll() []*Task
	Add(e *Task) (int, string)
	// Update(e *Task)
	// Delete(id int)
}

type taskService struct {
	sync.RWMutex
	m   map[int]*Task
	seq int
}

var tasks TaskService

func init() {
	tasks = &taskService{
		m: make(map[int]*Task),
	}

	tasks.Add(&Task{Id: 1, Script: "echo 'hello world!'"})
	println("Inserted one")
}

func (t *taskService) GetAll() []*Task {
	t.RLock()
	defer t.RUnlock()
	if len(t.m) == 0 {
		println("Length is 0")
		return nil
	}
	ret := make([]*Task, len(t.m))
	i := 0
	for _, v := range t.m {
		ret[i] = v
		i++
	}
	return ret
}

func (t *taskService) Get(id int) *Task {
	t.RLock()
	defer t.RUnlock()
	return t.m[id]
}

func (t *taskService) Add(e *Task) (int, string) {
	t.Lock()
	defer t.Unlock()

	if t.exists(e) {
		println("Can't add")
		return 0, "Already exists"
	}

	t.seq++
	e.Id = t.seq

	t.m[e.Id] = e
	return e.Id, ""
}

func (t taskService) exists(e *Task) bool {
	return false
}
