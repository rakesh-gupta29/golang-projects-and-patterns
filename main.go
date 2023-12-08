package main

type Updater interface {
	Update()
}

type Transform struct {
	x int32
}

type PlayerB struct {
	Transform
}
type PlayerA struct {
	Transform
}

func (player *PlayerA) Update() {
	player.x += 1

}

func main() {
	playerA := &PlayerA{}
	for i := 0; i < 100; i++ {
		playerA.Update()
	}

}

func Update(u Updater) {
	u.Update()
}
