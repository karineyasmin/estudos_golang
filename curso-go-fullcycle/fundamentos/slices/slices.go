package main

import "fmt"

func main() {
	s := []int{10, 20, 30, 50, 60, 40, 80, 90, 100}
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:0]), cap(s[:0]), s[:0])
	fmt.Printf("len=%d cap=%d %v\n", len(s[:4]), cap(s[:4]), s[:4])
	fmt.Printf("len=%d cap=%d %v\n", len(s[2:]), cap(s[2:]), s[2:])

	s = append(s, 110)
	fmt.Printf("len=%d cap=%d %v\n", len(s[:2]), cap(s[:2]), s[:2])
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

}

// cap -> capacidade
// [:x] -> tudo a direita desaparece
// [:0] -> apaga tudo
// [:x] -> apaga os x primeiros
// [x:] -> ignora os x primeiros
// toda vez que um append é feito, ele dobra o tamanho do slice
