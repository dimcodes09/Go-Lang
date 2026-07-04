package main

import "fmt"

func main(){
	var accountBalance  = 1000.0


	fmt.Println("Welcome to Go Bank")
	fmt.Println("What do you want to do? ")
	fmt.Println("1. Check Balance ")
	fmt.Println("2. Deposit MOney")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("YOur Choice: ")
	fmt.Scan(&choice)
	
	// wantsCheckBalance := choice == 1

	if choice == 1 {
		fmt.Println("Your balance is : ", accountBalance)
	} else if choice == 2{
		fmt.Print("Your deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		accountBalance += depositAmount //accountBalance = accountBalance + depositAMount
		fmt.Println("Balance updated! New Amount: " , accountBalance)
	} else if choice == 3{
		fmt.Print("Amount you want to withdraw: ")
	var withdrawmoney float64
	fmt.Scan(&withdrawmoney)
	accountBalance -= withdrawmoney
	fmt.Println("Balance UPdated ! New amount: ", accountBalance)
	}
}