package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error){
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 1000 , errors.New("FAIled to find Balance file")
	}

	balanceText := string(data)
	balance ,  err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return  1000 , errors.New("failed to parse stored balance value")
	}

	return balance , nil
}
func writebalanceToFile(balance float64){
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main(){
	var accountBalance , err  = getBalanceFromFile()

	if err != nil{
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("_____________")
		panic("can't continue , sorry")
	}

	fmt.Println("Welcome to Go Bank")

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
		writebalanceToFile(accountBalance)
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
	writebalanceToFile(accountBalance)
	default:
		fmt.Println("Goodbye! ")
		return

	}
  }
}


