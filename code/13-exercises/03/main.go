package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{vehicle: vehicle{doors: 4, color: "red"}, fourWheel: true}
	s := sedan{vehicle: vehicle{doors: 4, color: "blue"}, luxury: false}

	fmt.Printf("t: %v\n", t)
	fmt.Println(t.fourWheel)
	fmt.Printf("s: %v\n", s)
	fmt.Println(s.luxury)
}
