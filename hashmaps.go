package main

import "fmt"

func callerr() {
	hashMap := make(map[string]int)

	hashMap["apple"] = 5
	hashMap["banana"] = 3
	hashMap["orange"] = 7

	fmt.Println("apple:", hashMap["apple"])

	value, exists := hashMap["banana"]
	if exists {
		fmt.Println("banana:", value)
	} else {
		fmt.Println("banana not found")
	}

	delete(hashMap, "orange")

	for key, value := range hashMap {
		fmt.Printf("%s: %d\n", key, value)
	}
}
