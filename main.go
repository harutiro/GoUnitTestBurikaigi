package main

import "fmt"

type ClosedRange struct {
	ClosedRangeList []int
}

func NewClosedRange(
	lowerEndpoint int, 
	upperEndpoint int,
) *ClosedRange {
	closedRangeList := []int{}
	fmt.Println(closedRangeList)
	for i := lowerEndpoint; i <= upperEndpoint; i++ {
		closedRangeList = append(closedRangeList, i)
		fmt.Println(closedRangeList)
		fmt.Println(i)
	}
	return &ClosedRange{
		ClosedRangeList: closedRangeList,
	}
}

func (cr *ClosedRange) GetClosedList() []int{
	return cr.ClosedRangeList
}

func (cr *ClosedRange) GetLowerEndpoint() *int {
	if len(cr.ClosedRangeList) == 0 {
		return nil
	} else {
		return &cr.ClosedRangeList[0]
	}
}

func (cr *ClosedRange) GetUpperEndpoint() *int {
	if len(cr.ClosedRangeList) == 0 {
		return nil
	} else {
		return &cr.ClosedRangeList[len(cr.ClosedRangeList)-1]
	}
}


func (cr *ClosedRange) ToString() string {
	return fmt.Sprintf("[%+v,%+v]\n",*cr.GetLowerEndpoint(),*cr.GetUpperEndpoint())
}

func main() {
	cr := NewClosedRange(3,7)
	fmt.Println("Closed Range List:", cr.ClosedRangeList)
	fmt.Printf("Lower Endpoint: %+v\n", *cr.GetLowerEndpoint())
	fmt.Printf("Upper Endpoint: %+v\n", *cr.GetUpperEndpoint())
	fmt.Println("IntToString:", cr.ToString())
}
