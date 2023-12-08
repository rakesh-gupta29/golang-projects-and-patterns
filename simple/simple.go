package simple

import (
	"fmt"
	"math/rand"
)

type Man struct {
	age       int
	isMarried bool
}

func (man Man) isVoter() {
	fmt.Println(man.age > 18)
}

func CreateAndUseMen() {

	men := make([]Man, 100)
	for i := 0; i < 100; i++ {
		men[i] = Man{
			age:       rand.Intn(100),
			isMarried: rand.Intn(100) < 50,
		}
	}
	for i := 0; i < len(men); i++ {
		men[i].isVoter()
	}
}
