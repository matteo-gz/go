
```go
type iface struct {
	tab  *itab
	data unsafe.Pointer
}

type eface struct {
	_type *_type
	data  unsafe.Pointer
}
```



值为nil 类型不为nil x==nil =>false


编译器会由此检查 *myWriter 类型是否实现了 io.Writer 接口

`var _ io.Writer = (*myWriter)(nil)`