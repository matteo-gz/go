package main

import (
	"fmt"
	"unsafe"
)

type iface struct {
	itab, data uintptr
}

func main() {
	var a interface{} = nil
	// true
	fmt.Println(a == nil)

	var b interface{} = (*int)(nil)
	// false nil
	fmt.Println(b == nil, b)

	f := new(int)
	var e interface{} = &f
	//e  false 0x14000120020 0
	fmt.Println("e ", e == nil, e, *f)

	x := 5
	var c interface{} = (*int)(&x)

	// false 0x14000122020
	fmt.Println(c == nil, c)

	ia := *(*iface)(unsafe.Pointer(&a))
	ib := *(*iface)(unsafe.Pointer(&b))
	ic := *(*iface)(unsafe.Pointer(&c))

	//4329638912 0
	fmt.Println(ib.itab, ib.data)

	// {0 0} {4365306880 0} {4365306880 1374389997248}
	fmt.Println(ia, ib, ic)

	fmt.Println(*(*int)(unsafe.Pointer(ic.data)))
	// 5
}
