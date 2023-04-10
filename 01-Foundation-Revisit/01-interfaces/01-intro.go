package main

import "fmt"

func main() {
	var x interface{}
	x = 100
	x = "Culpa reprehenderit consectetur eu eu velit consequat reprehenderit id aliqua in. Anim culpa non elit reprehenderit ipsum proident aliquip cupidatat incididunt sit consequat. Quis cillum enim magna irure veniam amet sint nisi exercitation ipsum duis nulla. Ea esse deserunt ad aute id laborum duis reprehenderit reprehenderit reprehenderit dolor quis voluptate. Esse incididunt pariatur labore reprehenderit adipisicing non exercitation voluptate qui duis ea."
	x = true
	x = 99.99
	x = 14 + 15i
	x = struct{}{}
	fmt.Println(x)

	// x = 1000
	x = "Velit est ad ex sint esse laborum dolor eiusmod qui. Ex veniam enim deserunt culpa aliquip irure officia magna ullamco excepteur do incididunt Lorem. Qui adipisicing sint pariatur esse ullamco officia officia aute esse commodo sunt cillum exercitation qui. Esse in aliquip Lorem exercitation amet. Eu consectetur ullamco sit nulla voluptate elit pariatur nulla magna nisi sit aute ipsum sunt."
	// y := x.(int) + 2000

	// type assertion
	if val, ok := x.(int); ok {
		y := val + 2000
		fmt.Println(y)
	} else {
		fmt.Println("x is not an int")
	}
	// fmt.Println(y)

	// x = 1000
	// x = true
	x = 99.99
	switch val := x.(type) {
	case int:
		fmt.Println("x + 2000 = ", val+2000)
	case bool:
		fmt.Println("x is a bool, !x = ", !val)
	default:
		fmt.Println("unknown type")
	}

}
