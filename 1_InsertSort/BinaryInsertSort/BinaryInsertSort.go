package main

import (
	"fmt"
)

func main()  {
	var y = [...]float64{ 0.0, 6.0, 2.2, 7.2, 4.0, 9.0, 0.8, 5.0, 6.0, 4.5 }		//第一项为哨兵项，不参与排序
	
	fmt.Printf("排序前 : ")
	Output(y[:])

	BinaryInsertSort(y[:])

	fmt.Printf("排序后 : ")
	Output(y[:])
}


func Output(list []float64)  {
	for i := 1; i < len(list); i++ {
		fmt.Printf("%v  ", list[i])
	}
	fmt.Printf("\n")
}

func BinaryInsertSort(list []float64)  {
	var size int = len(list) - 1
	var i, j int
	var high, low, middle int
	for i = 2; i <= size; i++ {
		list[0] = list[i]
		high = i - 1;
		low = 1;
		for {
			if low > high {
				break;
			}
			middle = (high + low) / 2;
			if list[0] > list[middle] {
				low = middle + 1
			} else {
				high = middle - 1
			}
		}
		for j = i - 1; j >= high + 1; j-- {
			list[j + 1] = list[j]
		}
		list[high + 1] = list[0]
	}
}