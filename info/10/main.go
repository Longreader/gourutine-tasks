package main

import "fmt"

func update(p *int) {
	b := 2
	fmt.Println(b, "b")
	fmt.Println(&b)

	fmt.Println(p, "p")
	fmt.Println(*p)

	*p = b
}

func main() {
	var (
		a = 1
		p = &a
	)

	fmt.Println(a)
	fmt.Println(&a)

	fmt.Println(p)
	fmt.Println(*p)

	// fmt.Println(*p)
	update(p)
	// fmt.Println(*p)

	fmt.Println(p)
	fmt.Println(*p)

}
