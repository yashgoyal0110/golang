package main

import "fmt"

func main() {
	nums := []int{6, 7, 8}

	for i, num := range nums {
		fmt.Println(num, i)
	}

	m := map[string]string{"fname": "john", "lname": "doe"}

	for k, v := range m {
		fmt.Println(k, v)
		fmt.Println(k, v)
	}

	for k := range m {
		fmt.Println(k)
	}

	for i, c := range "golang" {
		fmt.Println(i, string(c))
	}
}
