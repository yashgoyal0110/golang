package main

func main() {
	// Creating a map
	m := make(map[string]string)

	// Setting elements
	m["name"] = "golang"
	m["area"] = "backend"

	// Getting elements
	// If key does not exist, returns zero value
	println(m["name"], m["area"])

	m2 := make(map[string]int)
	m2["age"] = 30
	m2["price"] = 50
	println(m2["phone"])
	println(len(m2))

	// Deleting an element
	delete(m2, "price")

	// Initializing map with values
	m3 := map[string]int{"price": 40, "phones": 3}

	// Checking if key exists
	v, ok := m3["phones"]
	println(v)
	if ok {
		println("all ok")
	} else {
		println("not ok")
	}

	m1 := map[string]int{"price": 40, "phones": 3}
	m4 := map[string]int{"price": 40, "phones": 8}

	println(m1 == nil)
	println(m4 == nil)

	m5 := m1
	m5["price"] = 100
	println(m1["price"])
	println(m5["price"])

	for key, value := range m1 {
		println(key, value)
	}
}
