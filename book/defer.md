[defer的陷阱 qcrao](https://qcrao.com/post/how-to-keep-off-trap-of-defer/)

> src/runtime/panic.go

`deferproc`

### defer 拆解
`return xxx`

> 1. 返回值 = xxx
> 2. 调用defer函数
> 3. 空的return

### 作用
- 释放资源
- 释放锁
- 记录日志
- 提高代码可读性
- 延迟调用,后进先出
- 与recover 可配合使用

```go
    defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover:", err)
		}
	}()

	fmt.Println("doing something")
	panic("oops! something went wrong")
```