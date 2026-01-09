package main

import "testing"
import "github.com/stretchr/testify/assert"

// 通常版
func TestClosedRange_Normal(t *testing.T) {
	var lowerNumber int = 3
	var upperNumber int = 8
	closedRangeList := []int{3,4,5,6,7,8}
	var containsCheck1 int = 5
	var containsCheck2 int = 2

	closedRange := NewClosedRange(lowerNumber,upperNumber)
	
	assert.Equal(t,*closedRange.GetLowerEndpoint(),lowerNumber)
	assert.Equal(t,*closedRange.GetUpperEndpoint(),upperNumber)
	assert.Equal(t,closedRange.GetClosedList(),closedRangeList)
	assert.True(t,closedRange.Contains(containsCheck1))
	assert.False(t,closedRange.Contains(containsCheck2))
}

// 負の値を含む範囲
func TestClosedRange_Negative(t *testing.T) {
	var lowerNumber int = -3
	var upperNumber int = 4
	closedRangeList := []int{-3,-2,-1,0,1,2,3,4}
	var containsCheck1 int = 0
	var containsCheck2 int = -5

	closedRange := NewClosedRange(lowerNumber,upperNumber)
	
	assert.Equal(t,*closedRange.GetLowerEndpoint(),lowerNumber)
	assert.Equal(t,*closedRange.GetUpperEndpoint(),upperNumber)
	assert.Equal(t,closedRange.GetClosedList(),closedRangeList)
	assert.True(t,closedRange.Contains(containsCheck1))
	assert.False(t,closedRange.Contains(containsCheck2))
}

// 上端と下端が同値
func TestClosedRange_Equal(t *testing.T) {
	var lowerNumber int = 3
	var upperNumber int = 3
	closedRangeList := []int{3}
	var containsCheck1 int = 3
	var containsCheck2 int = 2

	closedRange := NewClosedRange(lowerNumber,upperNumber)
	
	assert.Equal(t,*closedRange.GetLowerEndpoint(),lowerNumber)
	assert.Equal(t,*closedRange.GetUpperEndpoint(),upperNumber)
	assert.Equal(t,closedRange.GetClosedList(),closedRangeList)
	assert.True(t,closedRange.Contains(containsCheck1))
	assert.False(t,closedRange.Contains(containsCheck2))
}

// 上下が反転した場合
func TestClosedRange_Invers(t *testing.T) {
	var lowerNumber int = 7
	var upperNumber int = 3
	closedRangeList := []int{}
	var containsCheck1 int = 5
	var containsCheck2 int = 2

	closedRange := NewClosedRange(lowerNumber,upperNumber)

	assert.Empty(t,closedRange.GetLowerEndpoint())
	assert.Empty(t,closedRange.GetUpperEndpoint())
	assert.Equal(t,closedRange.GetClosedList(),closedRangeList)
	assert.False(t,closedRange.Contains(containsCheck1))
	assert.False(t,closedRange.Contains(containsCheck2))
}