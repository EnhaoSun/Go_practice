package main

import (
	"bytes"
	"fmt"
)

// tips:
/**
# 原理
假设要求11个数(0~10)存储在位数组里，11个数就需要11个bit位， 而1个byte有8个bit位，故须2个byte才能存储11个数。0存储在第0个byte的第0位，1存储在第0个byte的第1位...存储在第1个byte的第0位，9存储在第1个byte的第1位，10存储在第1个byte的第2位。如下

    |high                      low|
    |--- 1 byte ---|--- 0 byte ---|
    0 0 0 0 0 0 0 0|0 0 0 0 0 0 0 0
             10 9 8 7 6 5 4 3 2 1 0
一个数x(非负整数)若要存储在位数组里， 会面临两个问题：

x存储在第几个byte里？
因为一个byte可以存储8个bit，那很显然，x应该存储在第(x/8)个byte里。

x存储在第(x/8)个byte的第几位上？
通过观察，x应该存储在第(x%8)位上。

综上， x应该存储在第(x/8)个byte的第(x%8)位上。

在计算机里，位操作比除法和求模操作更高效点。x/8 相当于 x>>3(右移一位，相当于除以2；右移三位，相当于除以8)；x%8相当于x&0x7。

我们要使得x存储在位数组里，就是要使得第(x>>3)个byte的第(x&0x7)位上置1。由知识准备第3点的置1，可以得出，要使第0位置1，则与之相与的数为1(00000001)；要使第1位置1，则与之相与的数为2(00000010)；要使第2位置1，则与之相与的数为4(00000100) ...

由上可得，把x存储在位数组里，则需要：

bitArray[x>>3] |= (1<<(x&0x7))

进而可以得出，判断x存储在位数组里，则需要返回

bitArray[x>>3] & (1<<(x&0x7))

把x从位数组里删除，则需要

bitArray[x>>3] &= ~(1<<(x&0x7))

----------------------------------------------------------------------------
# 置1、置0，判断置位。

eg: x = 221,其二进制为11011101，若要将它的第5位由0变为1(置1)，则只须第5位与1或即可

    11011101
    00100000
   |--------
    11111101
若要将它的第5位由1变为0(置0)，则只须第4位与0与即可

    11011101
    11101111
   &--------
    11001101
判断第5位是否置1

    11111101
    00100000
   &--------
    00100000  (相与的结果非0则表明相应位已置1，否则置0)
----------------------------------------------------------------------------
 */

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x&0x3F)
	//word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x&0x3F)
	//word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	// turn the "bit"'th bit on the "word"th element to 1
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= t.words[i]
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	// 1 char = 1 byte(字节)
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}

		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
func main() {
	var x IntSet
	for i := 0; i < 100; i++  {
		x.Add(i)
	}
	fmt.Println(x.String())
	fmt.Println(x.Has(101))
}
