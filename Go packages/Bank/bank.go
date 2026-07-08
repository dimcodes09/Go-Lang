package main

import (
	"fmt"
	"github.com/dimcodes09/Go-Lang/GoPackages/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"


func main(){
	var accountBalance , err  = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil{
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("_____________")
		
	}

	fmt.Println("Welcome to Go Bank")
	fmt.Println("Reach Us 24x7", randomdata.PhoneNumber())

	for {

		presentOptions()

	var choice int
	fmt.Print("YOur Choice: ")
	fmt.Scan(&choice)
	
	// wantsCheckBalance := choice == 1

	switch choice {
	case 1:
		fmt.Println ("Your balance is : ", accountBalance)
	case 2:
		fmt.Print("Your deposit: ")
		var depositAmount float64
		fmt.Scan(&depositAmount)

		if depositAmount <= 0{
			fmt.Println("Invalid amt . Must be greater than 0.")
			// return
			continue
		}

		accountBalance += depositAmount //accountBalance = accountBalance + depositAMount
		fmt.Println("Balance updated! New Amount: " , accountBalance)
		fileops.WriteFloatToFile(accountBalance , accountBalanceFile)
	case 3:
		fmt.Print("Amount you want to withdraw: ")
	var withdrawmoney float64
	fmt.Scan(&withdrawmoney)

		if withdrawmoney <= 0{
			fmt.Println("Invalid amt . Must be greater than 0.")
			continue
		}

		if withdrawmoney > accountBalance{
			fmt.Println("Invalid amt . cant withdraw more than you have")
			// continue
		}

	accountBalance -= withdrawmoney
	fmt.Println("Balance UPdated ! New amount: ", accountBalance)
		fileops.WriteFloatToFile(accountBalance , accountBalanceFile)
	default:
		fmt.Println("Goodbye! ")
		return

	}
  }
}


