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
