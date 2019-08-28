package main

import (
	"fmt"
)

func main()  {
	var y = [...]float64{ 0.0, 6.0, 2.2, 7.2, 4.0, 9.0, 0.8, 5.0, 6.0, 4.5 }		//第一项为哨兵项，不参与排序
	
	fmt.Printf("排序前 : ")
	Output(y[:])

	SimpleSelectSort(y[:])

	fmt.Printf("排序后 : ")
	Output(y[:])
}


func Output(list []float64)  {
	for i := 1; i < len(list); i++ {
		fmt.Printf("%.2f  ", list[i])
	}
	fmt.Printf("\n")
}

func SimpleSelectSort(list []float64)  {
	var size int = len(list) - 1
	var i, j int
	var min int
	for i = 1; i <= size; i++ {
		min  = i
		for j = i + 1; j <= size; j++ {
			if list[j] < list[min] {
				min = j
			}
		}
		if(min != i) {
			/*
			a b交换:
				a = a + b
				b = a - b
				a = a - b
			*/
			list[i] = list[i] + list[min]
			list[min] = list[i] - list[min]
			list[i] = list[i] - list[min]
		}
	}
}