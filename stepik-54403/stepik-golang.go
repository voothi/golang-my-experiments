package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	fmt.Scan(&a) // считаем переменную 'a' с консоли
	fmt.Scan(&b) // считаем переменную 'b' с консоли
	a = a * a
	b = b * 2
	c = a + b
	fmt.Println(c)
	//2
	//2
	//8
}
