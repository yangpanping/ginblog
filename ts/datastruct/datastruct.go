package main

import "fmt"

//用于编写数据结构中的算法实现

//直接插入排序(升序)
func insertsort(data []int) {
	var temp int
	for i := 1; i < len(data); i++ {
		temp = data[i]
		if data[i]<data[i-1]{
			for j:=i-1;j>=0&&data[j]>temp ;j--  {
				data[j+1],data[j]=data[j],data[j+1]
			}
		}
	}
}

//折半插入排序（升序）
func insertsort2(data []int)  {
	var i,j,low,mid,high int
	for i=1;i<len(data) ;i++  {
		low,high=0,i-1
		for ;low<=high ;  {
			mid =(low+high)/2
			if data[mid]>data[i] {
				high=mid-1
			}else {
				low=mid+1
			}
		}
		for j=i-1;j>=high+1 ;j-- {
			data[j+1],data[j]=data[j],data[j+1]
		}

	}
}

//希尔排序
func shellsort(data []int)  {

}
func main()  {
	data :=[]int{49,38,65,97,76,13,27,49}

	//insertsort(data)
	insertsort2(data)
	fmt.Println(data)
}