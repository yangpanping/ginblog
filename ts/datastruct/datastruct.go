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

//堆排序
func HeapSortMax(arr []int, length int) []int {
	// length := len(arr)
	if length <= 1 {
		return arr
	}
	depth := length/2 - 1 //二叉树深度
	for i := depth; i >= 0; i-- {
		topmax := i //假定最大的位置就在i的位置
		leftchild := 2*i + 1
		rightchild := 2*i + 2
		if leftchild <= length-1 && arr[leftchild] > arr[topmax] { //防止越过界限
			topmax = leftchild
		}
		if rightchild <= length-1 && arr[rightchild] > arr[topmax] { //防止越过界限
			topmax = rightchild
		}
		if topmax != i {
			arr[i], arr[topmax] = arr[topmax], arr[i]
		}
	}
	return arr
}
func HeapSort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastlen := length - i
		HeapSortMax(arr, lastlen)
		if i < length {
			arr[0], arr[lastlen-1] = arr[lastlen-1], arr[0]
		}
	}
	return arr
}

func main() {
	data := []int{49, 38, 65, 97, 76, 13, 27, 49, 11, 23, 56, 97, 22, 54, 11, 45}

	//insertsort(data)
	//insertsort2(data)
	//bubbleSort(data)
	//quickSort(data, 0, len(data)-1)
	//selectSort(data)

	fmt.Println(HeapSort(data))
}
