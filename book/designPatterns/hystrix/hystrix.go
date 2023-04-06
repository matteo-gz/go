package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type CircuitBreaker struct {
	mu                  sync.Mutex
	isOpen              bool      // 断路器是否打开
	lastFailureTime     time.Time // 上一次失败时间
	consecutiveFailures int       // 连续失败次数
}

func (cb *CircuitBreaker) AllowRequest() bool {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	if cb.isOpen { // 如果断路器打开
		// 如果距离上次失败时间已经超过了 1 秒，那么重置断路器
		if time.Since(cb.lastFailureTime) > time.Second {
			cb.reset()
			return true
		} else {
			return false
		}
	} else { // 如果断路器关闭
		return true
	}
}

func (cb *CircuitBreaker) RecordFailure() {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	cb.consecutiveFailures++
	// 如果连续失败次数超过了 5 次，那么打开断路器
	if cb.consecutiveFailures >= 5 {
		cb.isOpen = true
		cb.lastFailureTime = time.Now()
	}
}

func (cb *CircuitBreaker) RecordSuccess() {
	cb.mu.Lock()
	defer cb.mu.Unlock()

	cb.reset()
}

func (cb *CircuitBreaker) reset() {
	cb.consecutiveFailures = 0
	cb.isOpen = false
}

func main() {
	cb := &CircuitBreaker{}
	client := &http.Client{}

	for i := 0; i < 20; i++ {
		// 检查是否允许发送请求
		if !cb.AllowRequest() {
			fmt.Println("Request blocked by circuit breaker")
			continue
		}

		// 发送 HTTP 请求
		resp, err := client.Get("http://localhost:8080/api")
		if err != nil {
			fmt.Println("Request failed with error:", err)
			cb.RecordFailure()
		} else {
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				fmt.Println("Unexpected status code:", resp.StatusCode)
				cb.RecordFailure()
			} else {
				fmt.Println("Request succeeded")
				cb.RecordSuccess()
			}
		}

		// 随机休眠一段时间
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	}
}
