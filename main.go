package main

import "fmt"

type ClosedRange struct {
	ClosedRangeList []int
}

func NewClosedRange(
	lowerEndpoint int, 
	upperEndpoint int,
) *ClosedRange {
	closedRangeList := make([]int, 0)
	for i := lowerEndpoint; i <= upperEndpoint; i++ {
		closedRangeList = append(closedRangeList, i)
	}
	return &ClosedRange{
		ClosedRangeList: closedRangeList,
	}
}

func (cr *ClosedRange) GetClosedList() []int{
	return cr.ClosedRangeList
}

func (cr *ClosedRange) GetLowerEndpoint() int {
	return cr.ClosedRangeList[0]
}

func (cr *ClosedRange) GetUpperEndpoint() int {
	return cr.ClosedRangeList[len(cr.ClosedRangeList)-1]
}

func main() {
	cr := NewClosedRange(3, 7)
	fmt.Println("Closed Range List:", cr.ClosedRangeList)
	fmt.Println("Lower Endpoint:", cr.GetLowerEndpoint())
	fmt.Println("Upper Endpoint:", cr.GetUpperEndpoint())
}



