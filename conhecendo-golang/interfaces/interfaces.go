package main

import "fmt"

type runner interface {
	run() string
	// turbo() string
}

// type turber interface {
// 	turbo() string
// }

// type turboRunning interface {
// 	runner
// 	turber
// }

type cat struct{}

func (c cat) run() string {
	return "cat is running"
}

type dog struct{}

func (d dog) run() string {
	return "dog is running"
}

type dogRobot struct {
	bateria int
	dog
}

// func (dr dogRobot) turbo() string {
// 	return "dogRobot is turbo"
// }

func run(r runner) {
	fmt.Println(r.run())
}

func main() {
	fuffy := cat{}
	thor := dog{}
	bot := dogRobot{}

	run(fuffy)
	run(thor)
	run(bot)
}
