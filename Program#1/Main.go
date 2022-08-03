package main

import "fmt"

func main() {
	const SMALLER_THEN_TWENTY int = 20
	const NOT_TEN int = 10
	var num1 int
	sum := 0
	for sum < SMALLER_THEN_TWENTY && sum != NOT_TEN {
		fmt.Println("Please Enter a number : ")
		fmt.Scan(&num1)
		sum += num1
		fmt.Printf("Your current sum is %d\n", sum)
	}
	fmt.Printf("Your total sum is : %v", sum)

}
