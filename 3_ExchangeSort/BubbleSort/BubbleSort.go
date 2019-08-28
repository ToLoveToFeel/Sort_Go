package main

import (
	"fmt"
)

func main()  {
	var y = [...]float64{ 0.0, 6.0, 2.2, 7.2, 4.0, 9.0, 0.8, 5.0, 6.0, 4.5 }		//第一项为哨兵项，不参与排序
	
	fmt.Printf("排序前 : ")
	Output(y[:])

	BubbleSort(y[:])

	fmt.Printf("排序后 : ")
	Output(y[:])
}


func Output(list []float64)  {
	for i := 1; i < len(list); i++ {
		fmt.Printf("%v  ", list[i])
	}
	fmt.Printf("\n")
}

func BubbleSort(list []float64)  {
	var size int = len(list) - 1
	var change_flag int	= 1 	//如果没发生交换说明已经排序完成，不需要继续运行，1代表发生变化

	for i := size; (i > 1) && (1 == change_flag); i-- {
		change_flag = 0
		for j := 1; j < i; j++ {
			if list[j] > list[j + 1] {
				list[0] = list[j];
				list[j] = list[j + 1];
				list[j + 1] = list[0];
				change_flag = 1;	//发生交换
			}
		}
	}
}