package pkg

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"sync"
)

func init() {
	// default max
	MAX_TASK_QUEUE = 20
	qNum := os.Getenv("MAX_TASK_QUEUE")
	if qNum != "" {
		num, err := strconv.Atoi(qNum)
		if err != nil {
			panic("fail get queue num")
		}
		MAX_TASK_QUEUE = num
	}
	TaskQueue = make(chan Task, MAX_TASK_QUEUE)
}

var MAX_TASK_QUEUE int

// task
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

// => GPM
var TaskQueue chan Task

type QuitSignal = struct{}
type ChanEmbedChanTask = chan chan Task

type Worker struct {
	WorkerPool  chan chan Task
	TaskChannel chan Task
	quit        chan QuitSignal
}

func NewWorker(wkPool ChanEmbedChanTask) Worker {
	return Worker{
		WorkerPool:  wkPool,
		TaskChannel: make(chan Task),
		quit:        make(chan QuitSignal),
	}
}

// real processer for task
func (w Worker) Start() {
	//var hungryNum int = 0
	for {
		// 1.1 first put task channel for dispatch can dispatch task
		// 1.1 work run early than dispatch
		// 2.put the channel to the worker pool when task execute finish
		w.WorkerPool <- w.TaskChannel

		select {
		case task := <-w.TaskChannel:
			task.Execute()
			fmt.Printf("worker result: %v\n", task.GetResult())
		case <-w.quit:
			fmt.Println("worker quit !")
			return
		}
	}
}

func (w Worker) Stop() {
	w.quit <- QuitSignal{}
}

// Dispach
type Dispatcher struct {
	WorkerPool ChanEmbedChanTask
	maxWorker  int
}

func NewDispatcher(maxWorkers int) *Dispatcher {
	pool := make(ChanEmbedChanTask, maxWorkers)
	return &Dispatcher{
		WorkerPool: pool,
		maxWorker:  maxWorkers,
	}
}

func (d *Dispatcher) Run() {

	for i := 0; i < d.maxWorker; i++ {
		worker := NewWorker(d.WorkerPool)
		go worker.Start()
	}

	go d.dispatch()
}

func (d *Dispatcher) dispatch() {

	//for task := range TaskQueue {
	for {
		if task, ok := <-TaskQueue; ok {
			go func(tsk Task) {
				taskChannel := <-d.WorkerPool

				taskChannel <- tsk
			}(task)
		} else {
			// need to quit
			fmt.Println("[TODO] need to quite !")
			return
		}

	}
}
