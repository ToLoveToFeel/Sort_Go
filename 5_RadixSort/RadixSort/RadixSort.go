package main

import (
	"fmt"
)

const MAX_NUM_OF_NUM int = 8	//关键字项数的最大值,此时整数最多8位
const NUM_OF_NUM int = 3	//实际关键字个数
const RADIX int = 10	//关键字的基数，此时是十进制整数的基数
const MAX_SPACE int = 100 

type SLCell struct {
	keys string
	next int
}

type SLList struct {
	r [MAX_SPACE]SLCell		//静态链表的可利用空间，r[0]为头结点
	keynum int		//记录的当前关键字（位数）个数
	recnum int		//静态链表的当前长度
}

func main()  {
	var y = [...]string{ "369", "367", "167", "239", "237", "138", "230", "139" }		
	var list SLList
	list.recnum = len(y)
	list.keynum = NUM_OF_NUM
	for i := 0; i < list.recnum; i++ {
		list.r[i].keys = y[i]
	}
	fmt.Printf("排序前 : ")
	for i := 0; i < len(y); i++ {
		fmt.Printf("%v  ", y[i])
	}

	RadixSort(&list);

	fmt.Printf("\n排序后 : ")
	RadixPrint(list);
}

func ord(keys string, i int) int {
	if i > len(keys) {
		return 0
	} else {
		return (int)(keys[i] - '0')
	}
}

func Distribute(R []SLCell, t int, f []int, r []int, head int) {
	var i int
	for i = 0; i < RADIX; i++ {
		f[i] = -1;	//链尾指针置为-1
		r[i] = -1;
	}
	for p := head; p != -1; p = R[p].next {
		i = ord(R[p].keys, NUM_OF_NUM - t - 1);		//ord将将记录中的第i个关键字映射到[]
		if (-1 == f[i]) {	//说明关键字为该数字的是第一个
			f[i] = p;
		} else {
			R[r[i]].next = p;
		}
		r[i] = p;
	}
}

func Collect(R []SLCell, f []int, r []int, head int) int {
	var i, t int
	for i = 0; i < RADIX && -1 == f[i]; i++ {

	}
	head = f[i]
	t = r[i]
	for {
		if !(i < RADIX) {
			break
		}
		i++
		for ; i < RADIX - 1 && -1 == f[i]; i++ {

		}
		if (i < RADIX && -1 != f[i]) {
			R[t].next = f[i];
			t = r[i];
		}
	}
	R[t].next = -1;

	return head
}

var head int = 0
func RadixSort(list *SLList)  {
	for i := 0; i < list.recnum; i++ {
		list.r[i].next = i + 1
	}
	list.r[list.recnum - 1].next = -1;
	var f = []int{ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1 }
	var r = []int{ -1, -1, -1, -1, -1, -1, -1, -1, -1, -1 }

	for i := 0; i < list.keynum; i++ {
		Distribute(list.r[:], i, f[:], r[:], head)
		head = Collect(list.r[:], f[:], r[:], head)
	}
}

func RadixPrint(list SLList)  {
	p := head
	for i := 0; i < list.recnum; i++ {
		fmt.Printf("%v  ", list.r[p].keys)
		p = list.r[p].next
	}
}
