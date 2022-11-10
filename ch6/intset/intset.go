// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Len return the number of elements
func (s *IntSet) Len() int64 {
	var length int64 = 0
	for _, world := range s.words {
		for i := uint64(1); i != 0; i <<= 1 {
			if (world & i) != 0 {
				length++
			}
		}
	}
	return length
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) bool {
	if x < 0 {
		return false
	}
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
	return true
}

func (s *IntSet) AddAll(nums ...int) bool {
	// 这里我们如果直接加进去的话，如果有负数的话会破坏之前的，而单独走一遍检查时间复杂度太高了
	// 这里我想到了log write，先写到一个临时的t中，然后求其与s的并集即可
	var t IntSet
	for _, i := range nums {
		if !t.Add(i) {
			return false
		}
	}
	s.UnionWith(&t)
	return true
}

// Remove remove x from the set
func (s *IntSet) Remove(x int) {
	if x < 0 {
		return
	}
	word, bit := x/64, uint(x%64)
	if word > len(s.words) {
		return
	}
	s.words[word] &= ^(1 << bit)
	// 更好的是如果如果是最后一个，那么释放掉
}

// Clear remove all elements from the set
func (s *IntSet) Clear() {
	s.words = nil

}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// Copy return a copy of the set
func (s *IntSet) Copy() *IntSet {
	var t IntSet
	t.words = make([]uint64, len(s.words))
	copy(t.words, s.words)
	return &t
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
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

//!-string
