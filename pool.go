package go_utils

import (
	"sync"
)

type Task struct {
	Id      int
	Job     func(t Task)
	Timeout int
}

type Pool struct {
	TaskQueue  chan Task
	Wg         sync.WaitGroup
	workersNum int
	bufferNum  int
}

func NewPool(workersNum int, chanBuffer int) *Pool {
	p := &Pool{
		workersNum: workersNum,
	}
	p.TaskQueue = make(chan Task, chanBuffer)
	return p
}

func (p *Pool) AddTask(task Task) {
	p.TaskQueue <- task
}

func (p *Pool) worker() {
	for task := range p.TaskQueue {
		task.Job(task)
	}
	p.Wg.Done()
}

func (p *Pool) GoTask(fn func()) {
	p.Wg.Add(1)
	go func() {
		fn()
		close(p.TaskQueue)
		p.Wg.Done()
	}()
}

func (p *Pool) Run() {
	p.Wg.Add(p.workersNum)
	for i := 0; i < p.workersNum; i++ {
		go p.worker()
	}
	p.Wg.Wait()
}
