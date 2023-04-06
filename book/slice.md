## slice

> [src/runtime/slice.go](https://github.com/matteo-gz/go/blob/note/1.19/src/runtime/slice.go)

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

## 练习题

### slice操作

> book/slice/print.go

```go



package main

import "fmt"

func main() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:6:7]

	s2 = append(s2, 100)
	s2 = append(s2, 200)

	s1[2] = 20

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(slice)
}
```
```
[2 3 20]
[4 5 6 7 100 200]
[0 1 2 3 20 5 6 7 100 9]
```

### 函数调用 改变索引

> book/slice/print2/print.go

```go

package main

import "fmt"

func main() {
	s := []int{1, 1, 1}
	f(s)
	fmt.Println(s)
}

func f(s []int) {
	// i只是一个副本，不能改变s中元素的值
	/*for _, i := range s {
	  	i++
	  }
	*/

	for i := range s {
		s[i] += 1
	}
}

```
```
[2 2 2]
```

### 调用指针与值

> book/slice/print3/print.go
```go
package main

import "fmt"

func myAppend(s []int) []int {
	// 这里 s 虽然改变了，但并不会影响外层函数的 s
	s = append(s, 100)
	return s
}

func myAppendPtr(s *[]int) {
	// 会改变外层 s 本身
	*s = append(*s, 100)
	return
}

func main() {
	s := []int{1, 1, 1}
	newS := myAppend(s)

	fmt.Println(s)
	fmt.Println(newS)

	s = newS

	myAppendPtr(&s)
	fmt.Println(s)
}

```
```
[1 1 1]
[1 1 1 100]
[1 1 1 100 100]
```

> 参考资料 https://golang.design/