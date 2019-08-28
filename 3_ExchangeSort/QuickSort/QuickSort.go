package main

import (
	"fmt"
)

func main()  {
	var y = [...]float64{ 0.0, 6.0, 2.2, 7.2, 4.0, 9.0, 0.8, 5.0, 6.0, 4.5 }		//第一项为哨兵项，不参与排序
	
	fmt.Printf("排序前 : ")
	Output(y[:])

	QuickSort(y[:], 1, len(y) - 1)

	fmt.Printf("排序后 : ")
	Output(y[:])
}


func Output(list []float64)  {
	for i := 1; i < len(list); i++ {
		fmt.Printf("%v  ", list[i])
	}
	fmt.Printf("\n")
}

func partition(list []float64, low int, high int) int {
	list[0] = list[low]
	for {
		if !(low < high) {
			break
		}
		for {
			if !((low < high) && (list[high] > list[0])) {
				break
			}
			high--
		}
		if low < high {
			list[low] = list[high]
			low++
		}

		for {
			if !((low < high) && (list[low] < list[0])) {
				break
			}
			low++
		}
		if (low < high) {
			list[high] = list[low];
			high--;
		}
	}
	list[high] = list[0];

	return high;
}

func QuickSort(list []float64, low int, high int)  {
	var t int
	if low < high {
		t = partition(list[:], low, high)
		QuickSort(list[:], low, t - 1)
		QuickSort(list[:], t + 1, high)
	}
}