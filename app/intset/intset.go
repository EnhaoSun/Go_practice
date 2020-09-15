package intset

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
若要将它的第5位由1变为0(置0)，则只须第5位与0与即可

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

//判断是32还是64
const wordSize = 32 << (^uint(0) >> 63)
const mask = wordSize - 1

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x&mask)
	//word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/wordSize, uint(x&mask)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	// turn the "bit"'th bit on the "word"th element to 1
	s.words[word] |= 1 << bit
}

// Add adds multiple the non-negative value xs to the set.
func (s *IntSet) AddAll(xs ...int) {
	for _, x := range xs {
		s.Add(x)
	}
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

		for j := 0; j < wordSize; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", wordSize*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

// Return set elements.
func (s *IntSet) Elems() []int {
	e := make([]int, 0)
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < wordSize; j++ {
			if word&(1<<uint(j)) != 0 {
				e = append(e, wordSize*i+j)
			}
		}
	}
	return e
}

// return the number of elements
// tips
/*
# 求二进制数中1的个数
快速法：运算次数与输入n的大小无关，只与n中1的个数有关。
如果n的二进制表示中有k个1，那么这个方法只需要循环k次即可。
其原理是不断清除n的二进制表示中最右边的1，同时累加计数器，直至n为0
```
int BitCount2(unsigned int n)
{
    unsigned int c =0 ;
    for (c =0; n; ++c)
    {
        n &= (n -1) ; // 清除最低位的1
    }
    return c ;
}
```
为什么n &= (n – 1)能清除最右边的1呢？
因为从二进制的角度讲，n相当于在n - 1的最低位加上1。
举个例子，8（1000）= 7（0111）+ 1（0001），
所以8 & 7 = （1000）&（0111）= 0（0000），清除了8最右边的1（其实就是最高位的1，因为8的二进制中只有一个1）。
再比如7（0111）= 6（0110）+ 1（0001），
所以7 & 6 = （0111）&（0110）= 6（0110），清除了7的二进制表示中最右边的1（也就是最低位的1）。
 */

func (s *IntSet) Len() int {
	num := 0
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; word > 0; j++ {
			word &= word - 1
			num++
		}
	}
	return num
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	word := x/64
	if word > len(s.words) {
		return
	}
	// turn the "bit"'th bit on the "word"th element to 0
	s.words[word] &= ^(1 << (x&mask))
}

//remove all elements from the set
func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	newset := &IntSet{}
	newset.words = make([]uint, len(s.words))
	copy(newset.words, s.words)
	return newset
}

// Set s to the intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= t.words[i]
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Set s to the difference of s and t. element in s not in t
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Set s to the symmetric difference of s and t.
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

/*
func main() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	fmt.Println(x.Len())
	x.AddAll(10, 11, 12)
	fmt.Println(x.String())
	x.words[0] = 0
	fmt.Println(x.String())
	fmt.Println(x.Elems())
}
 */