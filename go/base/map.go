package main

import (
	"fmt"
	"sort"
)

// 定义一个结构体
type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	// 1.创建map集合
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	// 2.map赋值操作
	personDB["1"] = PersonInfo{"1", "liutz", "beijing1"}
	personDB["2"] = PersonInfo{"2", "wuhc", "beijing2"}
	personDB["3"] = PersonInfo{"3", "liuyc", "beijing3"}

	// 3.删除一个key
	delete(personDB, "1")

	// 4.根据key查找map的Value
	person, ok := personDB["1"]

	if ok {
		fmt.Println("Found person", person.Name, "with ID 1")
	} else {
		fmt.Println("Did not find person with ID 1.")
	}

	// 5.map遍历key的顺序是随机的
	for k, v := range personDB {
		fmt.Println("k:", k, " v:", v)
	}

	// 6.制造一个有序key的map
	sortedKey := make([]string, 0)
	for k, _ := range personDB {
		sortedKey = append(sortedKey, k)
	}

	sort.Strings(sortedKey)

	for _, k := range sortedKey {
		fmt.Printf("k=%v,v=%v\n", k, personDB[k])
	}
}
