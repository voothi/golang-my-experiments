package main

import "fmt"

func main() {
	var d int
	fmt.Scan(&d)
	var time int
	var h int
	var m int
	time = d * 120
	h = time / 3600
	m = (time % 3600) / 60
	fmt.Println("It is", h, "hours", m, "minutes.")
}
