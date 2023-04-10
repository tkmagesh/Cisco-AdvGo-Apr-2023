package main

/*
	Note:
		Run the program with escape analysis
		go run -gcflags="-m" main.go
*/
import "fmt"

func main() {
	increment := getIncrement()
	fmt.Println(increment()) //=> 1
	fmt.Println(increment()) //=> 2
	// count = 10000            // influencing the outcome of the increment() function from outside
	fmt.Println(increment()) //=> 3
	fmt.Println(increment()) //=> 4
}

func getIncrement() func() int { //step-1
	var count int //step-2
	var name string = "Magesh"
	increment := func() int { //step-3
		count++ //step-4
		return count
	}
	fmt.Println(name)
	name = name + ""
	return increment //step-5
}
