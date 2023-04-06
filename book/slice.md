## slice

> https://github.com/matteo-gz/go/blob/note/1.19/src/runtime/slice.go

```go
type slice struct {
	array unsafe.Pointer // 指向底层数组的指针
	len   int            // 切片的长度
	cap   int            // 切片的容量
}

```

### 创建

`makeslice`

### 扩容

`func growslice(et *_type, old slice, cap int) slice `

growslice 函数用于在 append 操作中处理 slice 的扩容。
它的输入参数包括原 slice 的元素类型、旧的 slice、以及期望的新最小容量。
函数返回一个至少具有新容量的新 slice，其中包含旧数据的副本。新 slice 的长度被设置为旧 slice 的长度，而不是期望的新容量。
这是为了方便代码生成，因为旧 slice 的长度会立即用于计算在 append 操作期间写入新值的位置。

需要注意的是，在将来的版本中，growslice 函数可能会进行一些调整，例如返回新长度或仅返回指针/容量以节省堆栈空间等。但目前，旧 slice 的长度仍然被用于计算写入新值的位置。

总之，这段注释的目的是解释 growslice 函数的输入和输出参数以及其内部实现的一些细节，以便开发者更好地理解和使用该函数。


