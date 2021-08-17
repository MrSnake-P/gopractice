package wait

import (
	"sync"
	"time"
)

// 具有等待时间的互斥锁
type Wait struct {
	wg sync.WaitGroup
}

func (w *Wait) Add(delta int) {
	w.wg.Add(delta)
}

func (w *Wait) Done() {
	w.wg.Done()
}

func (w *Wait) Wait() {
	w.wg.Wait()
}

// 具有等待时间的阻塞，直到计数器为0或者是超时
func (w *Wait) WaitWithTimeOut(timeout time.Duration) bool {
	c := make(chan bool, 1)
	go func() {
		defer close(c)
		w.wg.Wait()
		c <- true
	}()
	select {
	case <-c:
		return false // 正常结束
	case <-time.After(timeout):
		return true	// 超时
	}
}

