package main

import "testing"
import "github.com/stretchr/testify/assert"

// 通常版
func TestClosedRange_Normal(t *testing.T) {
	var lowerNumber int = 3
	var upperNumber int = 8
	closedRangeList := []int{3,4,5,6,7,8}

	closedRange := NewClosedRange(lowerNumber,upperNumber)
	
	assert.Equal(t,closedRange.GetLowerEndpoint(),lowerNumber)
	assert.Equal(t,closedRange.GetUpperEndpoint(),upperNumber)
	assert.Equal(t,closedRange.GetClosedList(),closedRangeList)
}

// 負の値を含む範囲


// 0,0 の範囲


// 上下が反転した場合