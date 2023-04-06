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
