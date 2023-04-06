## 设计模式

- 责任链模式
- 熔断模式
- <a href="#section-3">函数选项模式</a>

## 责任链模式

> book/designPatterns/ChainofResponsibility/ChainofResponsibility.go

```go
package main

import "fmt"

// 定义处理器接口
type Handler interface {
	SetNext(Handler) Handler
	Handle(request int) (int, bool)
}

// 实现处理器结构体
type AddHandler struct {
	next Handler
}

func (a *AddHandler) SetNext(next Handler) Handler {
	a.next = next
	return next
}

func (a *AddHandler) Handle(request int) (int, bool) {
	if request < 0 {
		return 0, false
	}
	if request < 10 {
		return request, true
	}
	if a.next != nil {
		return a.next.Handle(request)
	}
	return 0, false
}

type DoubleHandler struct {
	next Handler
}

func (d *DoubleHandler) SetNext(next Handler) Handler {
	d.next = next
	return next
}

func (d *DoubleHandler) Handle(request int) (int, bool) {
	if request < 20 {
		return request * 2, true
	}
	if d.next != nil {
		return d.next.Handle(request)
	}
	return 0, false
}

func main() {
	// 初始化处理器链
	addHandler := &AddHandler{}
	doubleHandler := &DoubleHandler{}

	addHandler.SetNext(doubleHandler)

	// 使用处理器链处理请求
	result, ok := addHandler.Handle(5)
	if ok {
		fmt.Println("Result:", result)
	} else {
		fmt.Println("Failed to handle request")
	}

	result, ok = addHandler.Handle(15)
	if ok {
		fmt.Println("Result:", result)
	} else {
		fmt.Println("Failed to handle request")
	}

	result, ok = addHandler.Handle(-1)
	if ok {
		fmt.Println("Result:", result)
	} else {
		fmt.Println("Failed to handle request")
	}
}

```

##  熔断模式

> book/designPatterns/hystrix/hystrix.go

```go
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

```

## function option {#section-3}

> book/designPatterns/funcopt/funcopt.go

```go
package main

import "fmt"

type Options struct {
	Timeout int  // 超时时间（单位：秒）
	Debug   bool // 是否打印调试信息
}

type Option func(*Options)

func WithTimeout(timeout int) Option {
	return func(opts *Options) {
		opts.Timeout = timeout
	}
}

func WithDebug(debug bool) Option {
	return func(opts *Options) {
		opts.Debug = debug
	}
}

func DoSomethingWithOptions(opts ...Option) {
	// 设置默认值
	options := &Options{
		Timeout: 10,
		Debug:   false,
	}

	// 处理选项
	for _, opt := range opts {
		opt(options)
	}

	// 执行操作
	fmt.Printf("timeout=%d, debug=%v\n", options.Timeout, options.Debug)
}

func main() {
	// 不指定选项
	DoSomethingWithOptions()

	// 指定超时时间
	DoSomethingWithOptions(WithTimeout(30))

	// 指定是否打印调试信息
	DoSomethingWithOptions(WithDebug(true))

	// 指定所有选项
	DoSomethingWithOptions(WithTimeout(30), WithDebug(true))
}

```