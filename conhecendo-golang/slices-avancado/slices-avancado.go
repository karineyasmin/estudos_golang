package main

import "fmt"

// ... -> variadic function
func soma(a ...int) {
	total := 0
	for i := range a {
		total += a[i]
	}

	println(total)

}

func main() {
	var b []int = []int{1, 2, 3}
	b = append(b, 4)
	fmt.Println(b)

	// cap() -> capacidade do slice

	a := make([]int, 3, 5) // cria um slice com x elementos e x de capacidade
	fmt.Println(a)

	soma(1, 2, 3, 4, 5)
}
