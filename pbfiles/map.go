package main


import (
"fmt"
"sync"
)

func main() {
	// 创建一个 sync.Map 实例
	var myMap sync.Map
	var yourMap sync.Map
	yourMap.Store("provice","hebei")
	yourMap.Store("city","handan")
	yourMap.Store("code","057151")
	// 存储键值对
	myMap.Store("name", "Alice")
	myMap.Store("age", 30)
	provice,ok:=yourMap.Load("provice")
	if ok {
		fmt.Println("Provice",provice)
	}else {
		fmt.Println("Provice not found 404")
	}
	// 从 sync.Map 中加载值
	name, ok := myMap.Load("name")
	if ok {
		fmt.Println("Name:", name)
	} else {
		fmt.Println("Name not found")
	}

	age, ok := myMap.Load("age")
	if ok {
		fmt.Println("Age:", age)
	} else {
		fmt.Println("Age not found")
	}

	// 遍历 sync.Map 中的所有键值对
	myMap.Range(func(key, value interface{}) bool {
		fmt.Printf("%s: %v\n", key, value)
		return true // 返回 true 继续遍历，返回 false 停止遍历
	})

	// 删除键值对
	myMap.Delete("age")

	// 尝试加载已删除的键
	deletedAge, ok := myMap.Load("age")
	if !ok {
		fmt.Println("Age has been deleted")
	} else {
		fmt.Println("Deleted Age:", deletedAge)
	}

	// 加载或存储值（如果键已存在，则不更新值）
	loadedOrStored, loaded := myMap.LoadOrStore("name", "Bob")
	if loaded {
		fmt.Println("Name already exists, value:", loadedOrStored)
	} else {
		fmt.Println("Name loaded or stored, value:", loadedOrStored)
		// 注意：这里因为键 "name" 已经存在，所以不会存储 "Bob"，而是返回已存在的值 "Alice"
	}
}