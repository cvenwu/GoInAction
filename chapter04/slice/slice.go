package main

import "fmt"


//使用函数传递切片
func sliceDemo(nums []int) {
	nums [0] = 100  //修改可以起到作用
	//nums = append(nums, 1, 2, 3)   //添加元素不起作用
	//nums = nums[1:]  //切片也不起作用
	fmt.Println(nums)
}

//切片的内部实现

func main() {
	// 1. 切片的创建和初始化
	slice := make([]string, 5)  //创建一个长度为5的切片
	fmt.Printf("长度为----> %v, 容量为----> %v\n", len(slice), cap(slice))


	//2. 使用长度和容量声明切片
	slice2 := make([]int, 3, 5)
	fmt.Printf("长度为----> %v, 容量为----> %v\n", len(slice2), cap(slice2))

	//fmt.Println(slice2[3])


	//3. 声明一个nil切片
	var nums []int
	nums2 := make([]int, 0)
	nums3 := []int{}
	fmt.Println(nums, nums2)
	fmt.Println(nums == nil, nums2 == nil, nums3 == nil ) // true false false


	fmt.Println("修改前--------------------------------->")
	for index, val := range slice2 {
		fmt.Printf("index = %v, value = %v\n", index, val)
	}
	sliceDemo(slice2)
	fmt.Println("修改后--------------------------------->")
	for index, val := range slice2 {
		fmt.Printf("index = %v, value = %v\n", index, val)
	}
	fmt.Println(slice2)


}