package main

import (
	"fmt"
	"math"
	"strings"
	"time"
)

func greet() {
	fmt.Println("Hello, World!")
}

func add(a, b int) int {
	return a + b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func rectangle(length, width float64) (area, perimeter float64) {
	area = length * width
	perimeter = 2 * (length + width)
	return
}

func sum(numbers ...int) int {
	total := 0
	for _, num := range numbers {
		total += num
	}
	return total
}

func multiplier(factor int) func(int) int {
	return func(x int) int {
		return x * factor
	}
}

func filterEven(numbers []int) []int {
	var result []int
	for _, num := range numbers {
		if num%2 == 0 {
			result = append(result, num)
		}
	}
	return result
}

func wordCount(text string) map[string]int {
	words := strings.Fields(text)
	counts := make(map[string]int)
	for _, word := range words {
		counts[strings.ToLower(word)]++
	}
	return counts
}

type Person struct {
	Name string
	Age  int
}

func (p Person) introduce() string {
	return fmt.Sprintf("Hi, I'm %s and I'm %d years old", p.Name, p.Age)
}

func createPerson(name string, age int) Person {
	return Person{Name: name, Age: age}
}

func incrementAge(p *Person) {
	p.Age++
}

func factorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * factorial(n-1)
}

func demonstrateDefer() {
	defer fmt.Println("This prints last")
	fmt.Println("This prints first")
	fmt.Println("This prints second")
}

func asyncSum(numbers []int) int {
	ch := make(chan int)

	go func() {
		total := 0
		for _, num := range numbers {
			total += num
		}
		ch <- total
	}()

	return <-ch
}

func safeSqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, fmt.Errorf("cannot calculate square root of negative number: %v", x)
	}
	return math.Sqrt(x), nil
}

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func printArea(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
}

func main() {
	greet()

	fmt.Println("5 + 3 =", add(5, 3))

	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	area, perimeter := rectangle(5, 3)
	fmt.Printf("Rectangle: area=%.2f, perimeter=%.2f\n", area, perimeter)

	fmt.Println("Sum of 1,2,3,4,5:", sum(1, 2, 3, 4, 5))

	double := multiplier(2)
	fmt.Println("5 * 2 =", double(5))

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evens := filterEven(numbers)
	fmt.Println("Even numbers:", evens)

	text := "hello world hello go"
	counts := wordCount(text)
	fmt.Println("Word counts:", counts)

	person := createPerson("Alice", 25)
	fmt.Println(person.introduce())

	incrementAge(&person)
	fmt.Println("After birthday:", person.introduce())

	fmt.Println("Factorial of 5:", factorial(5))

	demonstrateDefer()

	fmt.Println("Async sum:", asyncSum([]int{1, 2, 3, 4, 5}))

	if sqrt, err := safeSqrt(16); err == nil {
		fmt.Println("Square root of 16:", sqrt)
	}

	circle := Circle{Radius: 5}
	printArea(circle)

	time.Sleep(time.Millisecond * 100)
}
