package main

import "sync"

type Task struct {
	Id     int    `json:"id"`
	Script string `json:"script"`
}

type TaskService interface {
	Get(id int) (*Task, bool)
	GetAll() []*Task
	Add(e *Task) (int, error)
	Update(e *Task) bool
	Delete(id int)
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
}

func (t *taskService) GetAll() []*Task {
	t.RLock()
	defer t.RUnlock()

	if len(t.m) == 0 {
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

func (t *taskService) Get(id int) (task *Task, ok bool) {
	t.RLock()
	defer t.RUnlock()

	task, ok = t.m[id]
	return
}

func (t *taskService) Add(e *Task) (int, error) {
	t.Lock()
	defer t.Unlock()

	t.seq++
	e.Id = t.seq
	t.m[e.Id] = e
	return e.Id, nil
}

func (t *taskService) Update(task *Task) bool {
	t.Lock()
	defer t.Unlock()

	if _, ok := t.m[task.Id]; !ok {
		return false
	}
	t.m[task.Id] = task
	return true
}

func (t *taskService) Delete(id int) {
	t.Lock()
	defer t.Unlock()
	delete(t.m, id)
}
