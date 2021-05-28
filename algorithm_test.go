package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func maxOnesAfterRemoveItem(arr []byte) uint {

	return 0
}

func TestSomething(t *testing.T) {
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
