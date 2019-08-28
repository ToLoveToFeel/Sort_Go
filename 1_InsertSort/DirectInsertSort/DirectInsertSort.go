package main

import (
	"fmt"
)


func main() {
	var y = [...]float64{ 0.0, 6.0, 2.2, 7.2, 4.0, 9.0, 0.8, 5.0, 6.0, 4.5 }		//第一项为哨兵项，不参与排序
	
	fmt.Printf("排序前 : ")
	Output(y[:])

	DirectInsertSort(y[:])

	fmt.Printf("排序后 : ")
	Output(y[:])
}

func Output(list []float64)  {
	for i := 1; i < len(list); i++ {
		fmt.Printf("%.2f  ", list[i])
	}
	fmt.Printf("\n")
}

func DirectInsertSort(list []float64)  {
	var i, j int 
	var size int = len(list) - 1
	for i = 2; i <= size; i++ {
		if list[i] < list[i - 1] {
			list[0] = list[i]
			list[i] = list[i - 1]
			for j = i - 2; list[j] > list[0]; j-- {
				list[j + 1] = list[j]
			}
			list[j + 1] = list[0]
		}
	} 
}