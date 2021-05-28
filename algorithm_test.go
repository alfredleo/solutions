package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func maxOnesAfterRemoveItem(arr []byte) uint {

	max, first, second, countZeroes := 0, 0, 0, 0
	for _, element := range arr {
		if element == 1 {
			first++
			countZeroes = 0
		} else {
			// we have one pair of first and second
			max = checkMax(first, second, max)
			// exchange first 1 count to second variable
			second = first
			first = 0

			countZeroes++
			if countZeroes > 1 {
				// reset counters if we have 0,0 or more
				first, second, countZeroes = 0, 0, 0
			}
		}
		//fmt.Println(first, second)
	}

	// fix corner case with small input or side 0s
	max = checkMax(first, second, max)

	// fix corner case with all 1s
	if len(arr) == max {
		max--
	}

	fmt.Println(max)
	return uint(max)
}

func checkMax(first int, second int, max int) int {
	sum := first + second
	if max < sum {
		max = sum
	}
	return max
}

func TestLongChainAlgorithm(t *testing.T) {
	assert := assert.New(t)

	assert.True(maxOnesAfterRemoveItem([]byte{0, 0}) == 0)
	assert.True(maxOnesAfterRemoveItem([]byte{0, 1}) == 1)
	assert.True(maxOnesAfterRemoveItem([]byte{1, 0}) == 1)
	assert.True(maxOnesAfterRemoveItem([]byte{1, 1}) == 1)
	assert.True(maxOnesAfterRemoveItem([]byte{1, 1, 0, 1, 1}) == 4)
	assert.True(maxOnesAfterRemoveItem([]byte{1, 1, 0, 1, 1, 0, 1, 1, 1}) == 5)
	assert.True(maxOnesAfterRemoveItem([]byte{1, 1, 0, 1, 1, 0, 1, 1, 1, 0}) == 5)
}


/*
I Задача о длиной цепочке единиц ⭐⭐

Дана последоватльность 0 и 1
Нужно найти саммую длинную последовательность из 1 (единиц) после удаления любого элемента

func maxOnesAfterRemoveItem([]byte) uint
assert(maxOnesAfterRemoveItem[0,0] == 0)
assert(maxOnesAfterRemoveItem[0,1] == 1)
assert(maxOnesAfterRemoveItem[1,0] == 1)
assert(maxOnesAfterRemoveItem[1,1] == 1)
assert(maxOnesAfterRemoveItem[1, 1, 0, 1, 1] == 4)
assert(maxOnesAfterRemoveItem[1, 1, 0, 1, 1, 0, 1, 1, 1] == 5)
assert(maxOnesAfterRemoveItem[1, 1, 0, 1, 1, 0, 1, 1, 1, 0] == 5)
Что хочется увидеть:

Алгоритм со сложностью O(N) по времени и O(1) по памяти
*/
