package main

import "fmt"


func demp01(mymap map[string]string) {

	if _, exist := mymap["red"]; exist {
		delete(mymap, "red")
	}
}

func main() {


	//1. 映射的创建和初始化的两种方式
	//1.1 使用make内置函数
	dict := make(map[string]int)   //创建一个映射，键的类型是string，值的类型是int
	//1.2 使用字面量
	dict2 := map[string]string{"0":"1", "1":"0"}  //使用两个键值对初始化映射

	fmt.Println(dict, dict2)


	//使用切片作为映射的键,报错
	//mymap := map[[]string]string{}

	//声明一个未经初始化的映射
	//var mymap map[string]int
	//mymap["123"] = 2  //报错 panic: assignment to entry in nil map


	colors := map[string]string{"red": "blue", "green":"yellow", "1": "2"}
	//val 将会取得该键对应的值，如果键不存在，将会取得值对应类型的零值，
	//exist判断键是否存在，如果存在将会返回true
	val, exist := colors["red"]
	if exist {
		fmt.Println(val)
	}

	//上面的26-29行代码等价于
	if val, exist := colors["red"]; exist {
		fmt.Println(val)
	}

	for key, val := range colors {
		fmt.Printf("key = %s, value = %s\n", key, val)
	}

	//从colors这个map中删除键为red的键值对
	delete(colors, "red")
	fmt.Println("删除后--------------------------------->")
	for key, val := range colors {
		fmt.Printf("key = %s, value = %s\n", key, val)
	}



	//使用函数进行删除
	colors2 := map[string]string{"red": "blue", "green":"yellow", "1": "2"}
	fmt.Println("删除前--------------------------------->")
	for key, val := range colors2 {
		fmt.Printf("key = %s, value = %s\n", key, val)
	}
	demp01(colors2)
	fmt.Println("删除后--------------------------------->")
	for key, val := range colors2 {
		fmt.Printf("key = %s, value = %s\n", key, val)
	}

}