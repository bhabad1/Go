package main

import "fmt"

func main(){
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you wanted to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposite Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	fmt.Print("Your Choice: ")
	var choice int
	fmt.Scan(&choice)
	fmt.Println("Your Choice: ", choice)
}
