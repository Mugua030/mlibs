package pkg

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	var wg sync.WaitGroup
	pool := NewPool(3)

	// productTask
	product := &ProductTask{
		Wg: &wg,
	}
	wg.Add(1)
	pool.Submit(product)

	// coder task
	code := &CoderTask{
		Wg: &wg,
	}
	wg.Add(1)
	pool.Submit(code)

	wg.Wait()

	// out
	fmt.Println(product.GetResult())
	fmt.Println(code.GetResult())
}

func TestTaskPool(t *testing.T) {
	go func() {
		for i := 0; i <= 18; i++ {
			tsk := NewBigDataTask(i)
			TaskQueue <- tsk
		}
	}()

	dspch := NewDispatcher(3)
	dspch.Run()

	select {}
	//fmt.Println("task finished !")
}

// for test task
type BigDataTask struct {
	Id     int
	Name   string
	Result string
}

func NewBigDataTask(id int) *BigDataTask {
	return &BigDataTask{
		Id:     id,
		Name:   "",
		Result: "",
	}
}

func (b *BigDataTask) Execute() {
	//fmt.Println("BigDataTask execute Done")
	time.Sleep(time.Second * 1)
	ret := fmt.Sprintf("BigDataTask[%d] Execute Done", b.Id)
	b.Result = ret
}
func (b *BigDataTask) GetResult() string {
	return b.Result
}
