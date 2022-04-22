package main

import (
	"fmt"
)

//用于编写数据结构中的算法实现

//直接插入排序(升序)
func insertsort(data []int) {
	var temp int
	for i := 1; i < len(data); i++ {
		temp = data[i]
		if data[i] < data[i-1] {
			for j := i - 1; j >= 0 && data[j] > temp; j-- {
				data[j+1], data[j] = data[j], data[j+1]
			}
		}
	}
}

//折半插入排序（升序）
func insertsort2(data []int) {
	var i, j, low, mid, high int
	for i = 1; i < len(data); i++ {
		low, high = 0, i-1
		for low <= high {
			mid = (low + high) / 2
			if data[mid] > data[i] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		}
		for j = i - 1; j >= high+1; j-- {
			data[j+1], data[j] = data[j], data[j+1]
		}

	}
}

//希尔排序
func shellsort(data []int) {

}

//冒泡排序
func bubbleSort(data []int) {
	for i := 0; i < len(data); i++ {

		for j := len(data) - 1; j > i; j-- {
			if data[j-1] > data[j] {
				data[j-1], data[j] = data[j], data[j-1]
			}

		}

	}
}

//快速排序
func quickSort(datas []int, low, high int) {
	if low < high {
		pivotpos := partition(datas, low, high)
		quickSort(datas, low, pivotpos-1)
		quickSort(datas, pivotpos+1, high)
	}
}
func partition(data []int, low, hegh int) int {
	temp := data[low]
	for low < hegh {
		for ; low < hegh && data[hegh] >= temp; hegh-- {
		}
		data[low] = data[hegh]
		for ; low < hegh && data[low] <= temp; low++ {
		}
		data[hegh] = data[low]
	}
	data[low] = temp
	return low
}

//简单选择排序
func selectSort(datas []int) {
	for i := 0; i < len(datas)-1; i++ {
		min := i
		for j := i + 1; j < len(datas); j++ {
			if datas[j] < datas[min] {
				min = j
			}
		}
		if min != i {
			datas[i], datas[min] = datas[min], datas[i]
		}
	}
}

func main() {
	data := []int{49, 38, 65, 97, 76, 13, 27, 49, 11, 23, 56, 97, 22, 54, 11, 45}

	//insertsort(data)
	//insertsort2(data)
	//bubbleSort(data)
	//quickSort(data, 0, len(data)-1)
	selectSort(data)
	fmt.Println(data)
}
