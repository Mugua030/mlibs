package pkg

import (
	"fmt"
	"sync"
	"testing"
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
