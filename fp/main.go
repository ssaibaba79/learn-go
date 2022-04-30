package main

import "fmt"

type point struct {
	x, y int
}

func main() {

	x := 20
	v(x)
	fmt.Printf("%p, %v\n", &x, x)

	y := []int{1, 2, 3}
	r(y)
	fmt.Printf("%p, %v\n", &y, y)

	c := [3]rune{'a', 'b', 'c'}
	v2(c)
	fmt.Printf("%p, %v\n", &c, c)

	m := make(map[string]int)
	r2(m)
	fmt.Printf("%p, %v\n", &m, m)

	s := point{}
	r3(s)
	fmt.Printf("%p, %v\n", &s, s)

	r4(&s)
	fmt.Printf("%p, %v\n", &s, s)
}

func v(a int) {

	a++
	fmt.Printf("%p, %v\n", &a, a)
}

func v2(c [3]rune) {

	c[0] = 'd'
	fmt.Printf("%p, %v\n", &c, c)
}

func r(a []int) {

	a[0]++
	fmt.Printf("%p, %v\n", &a, a)
}

func r2(m map[string]int) {

	m["one"]++
	fmt.Printf("%p, %v\n", &m, m)
}

func r3(s point) {
	s.x = 10
	fmt.Printf("%p, %v\n", &s, s)
}

func r4(s *point) {
	s.x = 10
	fmt.Printf("%p, %v\n", s, s)
}
