package main

import (
	"fmt"
)

func main()  {
	var y = [...]float64{ 0.0, 6.0, 2.2, 7.2, 4.0, 9.0, 0.8, 5.0, 6.0, 4.5 }		//第一项为哨兵项，不参与排序
	// var a [3]int
	var a = [...]int{5, 3, 1}
	
	fmt.Printf("排序前 : ")
	Output(y[:])

	for i := 0; i < len(a); i++{
		ShellSort(y[:], a[i])		
	}

	fmt.Printf("排序后 : ")
	Output(y[:])
}


func Output(list []float64)  {
	for i := 1; i < len(list); i++ {
		fmt.Printf("%.2f  ", list[i])
	}
	fmt.Printf("\n")
}

func ShellSort(list []float64, d int)  {
	var size int = len(list) - 1
	var i, j int
	for i = d + 1; i <= size; i++ {
		if list[i] < list[i - d] {
			list[0] = list[i]
			list[i] = list[i - d]
			for j = i - 2 * d; (j > 0) && (list[0] < list[j]); j = j - d {
				list[j + d] = list[j]
			}
			list[j + d] = list[0]
		}
	}
}