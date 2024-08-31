package pkg

import (
	"runtime"
	"sync"
)

type Task interface {
	Execute()
	GetResult() string
}

type RoutinePool struct {
	Num      int
	TaskList chan Task
}

func NewPool(num int) *RoutinePool {
	if num == 0 {
		num = runtime.NumCPU()
	}
	p := &RoutinePool{
		TaskList: make(chan Task),
	}

	// The working routine
	for i := 0; i < num; i++ {
		go func() {
			for task := range p.TaskList {
				task.Execute()
			}
		}()
	}

	return p
}

// submit task
func (r *RoutinePool) Submit(t Task) {
	r.TaskList <- t
}

// here: just for test
type ProductTask struct {
	Wg     *sync.WaitGroup
	Result string
}

func (p *ProductTask) Execute() {
	p.Result = "Producter: Design Future"
	p.Wg.Done()
}
func (p *ProductTask) GetResult() string {
	return p.Result
}

type CoderTask struct {
	Wg     *sync.WaitGroup
	Result string
}

func (c *CoderTask) Execute() {
	c.Result = "Coder: coding"
	c.Wg.Done()
}
func (c *CoderTask) GetResult() string {
	return c.Result
}

type TesterTask struct {
	Wg     *sync.WaitGroup
	Result string
}

func (t *TesterTask) Execute() {
	t.Result = "Tester: test"
	t.Wg.Done()
}

func (t *TesterTask) GetResult() string {
	return t.Result
}
