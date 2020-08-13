package main

import "fmt"

func main() {
	//声明一个数组并且设置为零值
	var nums [5]int
	fmt.Println(nums)

	//使用数组字面量来声明数组
	array := [5]int{1, 2, 3}
	arrays := [5]int{1, 2, 3}
	fmt.Println(array, arrays)

	//让Go自动计算数组的长度
	//容量由初始值的数量决定
	array2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array2, len(array2))

	//声明数组并指定特定元素的值
	array3 := [5]int{1: 10, 2: 20}
	fmt.Println(array3)

	//访问指针数组的元素
	array4 := [5]*int{0: new(int), 1: new(int)}
	*array4[0] = 100  //相当于*(array4[0])
	*(array4[1]) = 200
	fmt.Println(array4, *array4[0])
}
