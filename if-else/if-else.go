package main

import "fmt"

func main() {
	var num int = 24

	if num%2 != 0 {
		fmt.Println("Weird")
	} else {
		if num >= 2 && num <= 5 {
			fmt.Println("Not Weird")
		} else if num >= 6 && num <= 20 {
			fmt.Println("Weird")
		} else {
			fmt.Println("Not Weird")
		}

	}
}