
## 位操作

https://bit-calculator.com/bit-shift-calculator

### 按位与（&）：
对于每一位，只有两个操作数在该位上的值都为 1 时，结果的该位上的值才为 1，否则结果的该位上的值为 0。

```
    00101100
  & 11001011
  = 00001000

```

### 按位或（|）：
对于每一位，只要两个操作数在该位上的值中至少有一个为 1，结果的该位上的值就为 1，否则结果的该位上的值为 0。
```
    00101100
  | 11001011
  = 11101111

```
### 按位异或（^）：
对于每一位，当两个操作数在该位上的值不相同时，结果的该位上的值为 1，否则结果的该位上的值为 0。
```
    00101100
  ^ 11001011
  = 11100111

```
### 左移（<<）：
将一个二进制数向左移动指定的位数，相当于在二进制数的末尾添加指定数量的 0。
```
00101100 << 2 = 10110000

```
### 右移（>>）：
将一个二进制数向右移动指定的位数，相当于去掉二进制数末尾的指定数量的位数。
```
00101100 >> 2 = 1011

```
### map中的多标记写入
这些位运算符在 Go 语言中主要用于处理位标志和位掩码。
位标志是一种用于表示某个状态或属性的二进制标记，而位掩码则是一种用于提取、设置或清除位标志的掩码值。
在处理二进制数据时，这些运算符可以用于对数据的二进制表示进行精细的操作和处理。

```

10进制	2进制
0	0
1	1
2	10
3	11
4	100
5	101
6	110
7	111
8	1000
9	1001
10	1010
11	1011
12	1100
13	1101
14	1110
15	1111
16	10000
```



```go
// flags
// 1
iterator = 1 // there may be an iterator using buckets
// 10
oldIterator = 2 // there may be an iterator using oldbuckets
// 100
hashWriting = 4 // a goroutine is writing to the map
// 1000
sameSizeGrow = 8 // the current map growth is to a new map of the same size
```
```go


h.flags |= hashWriting

/*
假设
hashWriting=0b100
h.flags=0b1011
h.flags |= hashWriting
(h.flags==0b1111) ==true
*/
```
```
flags := h.flags &^ (iterator | oldIterator)
/*
假设
h.flags=0b1011
flags := h.flags &^ (iterator | oldIterator)
flags := 0b1011 &^ 0b11

 0b1011
 0b0011  //  bit 位为 1取0,为0取上面的值
&^=
 0b1000
 
 该操作清除了iterator和oldIterator 使之为0
*/
```