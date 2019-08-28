package main

import (
	"fmt"
)

func main()  {
	var y = [...]float64{ 0.0, 6.0, 2.2, 7.2, 4.0, 9.0, 0.8, 5.0, 6.0, 4.5 }		//第一项为哨兵项，不参与排序
	
	fmt.Printf("排序前 : ")
	Output(y[:])

	HeapSort(y[:])

	fmt.Printf("排序后 : ")
	Output(y[:])
}


func Output(list []float64)  {
	for i := 1; i < len(list); i++ {
		fmt.Printf("%v  ", list[i])
	}
	fmt.Printf("\n")
}

func HeapAdjust(list []float64, m int, n int)  {
	var i int
	list[0] = list[m]
	for i = 2 * m; i <= n; i *= 2 {
		if i < n && list[i] < list[i + 1] {
			i++
		}
		if list[0] > list[i] {
			break
		}
		list[m] = list[i]
		m = i
	}
	list[m] = list[0]
}

func HeapSort(list []float64)  {
	var size = len(list) - 1
	var i int
	for i = size / 2; i > 0; i-- {
		HeapAdjust(list[:], i, size)
	}
	for i = size; i > 1; i-- {
		list[0] = list[i]
		list[i] = list[1]
		list[1] = list[0]
		HeapAdjust(list[:], 1, i -1)
	}
}