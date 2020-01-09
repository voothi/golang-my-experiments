package main

import "fmt"

func main() {
	const a int = 5
	if a > 0 {
		fmt.Println("Число положительное")
	} else if a < 0 {
		fmt.Println("Число отрицательное")
	} else {
		fmt.Println("Ноль")
	}
}
