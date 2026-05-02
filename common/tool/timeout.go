package tool

import (
	"sync"
	"time"
)

// WaitTimeout
//
//	@Description: 多线程超时控制
//	@param wg 线程组
//	@param timeout 超时时间
//	@return bool false 未超时 true 超时
func WaitTimeout(wg *sync.WaitGroup, timeout time.Duration) bool {
	c := make(chan struct{})
	go func() {
		defer close(c)
		wg.Wait()
	}()
	select {
	case <-c:
		return false // completed normally
	case <-time.After(timeout):
		return true // timed out
	}
}
