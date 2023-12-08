package interfaces

import "fmt"

type Man struct {
	age int
}

func main() {
	man := Man{
		age: 10,
	}

	fmt.Print(man)
}
