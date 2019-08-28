package main

import (
	"fmt"
)

func main()  {
	var y = [...]float64{ 0.0, 6.0, 2.2, 7.2, 4.0, 9.0, 0.8, 5.0, 6.0, 4.5 }		//第一项为哨兵项，不参与排序
	
	fmt.Printf("排序前 : ")
	Output(y[:])

	MergeSort(y[:])

	fmt.Printf("排序后 : ")
	Output(y[:])
}


func Output(list []float64)  {
	for i := 1; i < len(list); i++ {
		fmt.Printf("%v  ", list[i])
	}
	fmt.Printf("\n")
}

func Merge(list2 []float64, list []float64, left int, middle int, right int)  {
	//list为合并后的目标数组，将有序的序列 list2[s..m] 和 list2[m+1..t]归并为有序的序列 list[s..t]
	var i, j, k int = left, middle + 1, left
	for ; (i <= middle) && (j <= right); k++ {
		if (list2[i] <= list2[j]) {
			list[k] = list2[i]
			i++
		} else {
			list[k] = list2[j]
			j++
		}
	}
	if (i <= middle) {
		for ; i <= middle; {
			list[k] = list2[i]
			k++
			i++
		}
	}
	if (j <= right) {
		for ; j <= right; {
			list[k] = list2[j]
			k++
			j++
		}
	}
}

func Msort(list2 []float64, list []float64, start int, end int)  {
	//将 list2[s..t] 通过 list3[] 归并排序为 list[s..t]
	var list3 = [...]float64{ 0.0, 6.0, 2.2, 7.2, 4.0, 9.0, 0.8, 5.0, 6.0, 4.5 }
	if (start == end) {
		list[start] = list2[start];
	} else {
		var middle int = (start + end) / 2;
		Msort(list2[:], list3[:], start, middle);		//体现了分久必合的思想
		Msort(list2[:], list3[:], middle + 1, end);
		Merge(list3[:], list[:], start, middle, end);
	}
}

func MergeSort(list []float64)  {
	var size int = len(list) - 1
	Msort(list[:], list[:], 1, size)
}