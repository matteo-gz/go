参考文章



## 逃逸分析
[Golang 之变量去哪儿](https://qcrao.com/post/where-go-variables-go/)

逃逸分析（Escape Analysis）是指编译器在编译时分析程序中的变量在函数调用栈中的生命周期，进而确定变量应该存储在栈上还是堆上的过程。

在函数内部声明的变量通常会被分配在栈上，这样可以避免因为动态内存分配而带来的性能开销。但是，有些变量可能需要在函数调用结束后继续存在，比如函数返回一个指向该变量的指针，或者该变量被赋值给了全局变量等。这些情况下，变量需要被分配在堆上，因为栈上的变量随着函数调用的结束而销毁，无法满足需求。

逃逸分析的作用是在编译时分析程序中的变量是否可能逃逸到函数外部，如果变量不会逃逸到函数外部，那么它可以安全地被分配在栈上，从而避免堆上内存分配和垃圾回收的开销。如果变量可能会逃逸到函数外部，那么它就需要在堆上分配内存。

逃逸分析是一项重要的优化技术，在 Go 语言中尤其重要，因为 Go 语言中的并发模型需要大量的内存分配和垃圾回收。通过逃逸分析，可以大大减少堆内存分配和垃圾回收的次数，从而提高程序的性能。

### 分析逃逸
`go build -gcflags '-m -l' main.go `

编译 Go 程序，其中 -gcflags '-m -l' 是指定编译器选项。

-gcflags 用于设置编译器的选项，其中 -m 用于进行内存分析，可以查看程序中哪些变量分配在栈上，哪些分配在堆上；
-l 则用于关闭一些无用代码的优化，例如函数内的未使用变量、未使用的导入包等，这样可以减小程序的体积。

### 查看汇编

`go tool compile -S main.go`

将 Go 源代码编译为汇编代码，以便于查看程序的底层实现。

-S 选项用于生成汇编代码，并输出到标准输出流。执行该命令后，会将 main.go 文件编译为汇编代码，并将汇编代码输出到控制台。
汇编代码可以让开发者更加深入地了解 Go 程序的底层实现细节，有助于进行性能调优和程序优化。

### 结论

如果一个变量的生命周期不会逃逸到函数外部，那么可以将其分配在栈上，否则必须分配在堆上。
函数外部的变量总是在堆上分配。

## 内存布局
[图解 Go 语言内存分配](https://qcrao.com/post/graphic-go-memory-allocation/)


线程内存池,全局内存池

> src/runtime/malloc.go

### heap组成

堆内存由多个arena组成
arenaBaseOffset是arena的起始地址

arena包含多个page,每个page大小为8KB

内存多级管理,降低锁粒度。算法起源是tcmalloc(内存碎片避免，需要将内存划分为多个规格)

arena划分多个块(`span`),不同规格的内存块放入链表中,规格有`_NumSizeClasses`种预置

> src/runtime/sizeclasses.go

详情看 `class_to_size`
```
heap=[[arena][arena][arena][...]]
arena=[[span][span]...]
span=[[page][page]...]
page=[[内存块][内存块]...]
```

### heap结构体

`mheap` 管理整个堆内存

`arena`为`headArena`

`span`为`mspan`

p.mcache 不用加锁

### local p需要mspan时
```


p.mcache.alloc查找

如果没有 或者用完了

就去mcentral这里获取一个mspan,
假设为 *mspan=p.mcache.alloc[x]

mcentral{
    partical [][]
    full     []
}

partical里的取出一个 与 p.mcache.alloc[x] 与之交换

after:

mcentral{
partical []
full     []
}

*mspan=mcentral.partical[new]

将 p.mcache.alloc[x]  之前用完的 放入 full

final:

mcentral{
partical []
full     [][]
}
```
### TLS变量

Thread Local Storage

`go:linkname` 

### 全局变量

allgs