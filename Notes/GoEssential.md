# Investment_caluclator code where you will learn how the values and types works in go 

package main

import ("math"
"fmt"
)

func main(){
	var investmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years float64 = 10

	var futureValue = float64(investmentAmount )* math.Pow( 1 + expectedReturnRate/100 , float64(years))
	fmt.Println(futureValue)
}

or 

# Alternative using ( := )

package main

import ("math"
"fmt"
)

func main(){
     investmentAmount := 1000.0 
	 years := 10.0
	 expectedReturnRate := 5.5
	 
	var futureValue = float64(investmentAmount )* math.Pow( 1 + expectedReturnRate/100 , float64(years))
	fmt.Println(futureValue)
}

# LEarning fmt.Scan()

package main

import ("math"
"fmt"
)

func main(){
	const inflationRate  = 6.5
    var  investmentAmount float64
	 var years float64
	 expectedReturnRate := 5.5

	 fmt.Print("Investment Amount: ")
	 fmt.Scan(&investmentAmount)

	 fmt.Print("Expected return rate: ")
	 fmt.Scan(&expectedReturnRate)

	 fmt.Print("Years: ")
	 fmt.Scan(&years)
	 
     futureValue := float64(investmentAmount )* math.Pow( 1 + expectedReturnRate/100 , years)
	 futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}


# OR

package main

import ("math"
"fmt"
)

func main(){
	const inflationRate  = 6.5
    var  investmentAmount float64
	 var years float64
	 expectedReturnRate := 5.5

	 fmt.Print("Investment Amount: ")
	 fmt.Scan(&investmentAmount)

	 fmt.Print("Expected return rate: ")
	 fmt.Scan(&expectedReturnRate)

	 fmt.Print("Years: ")
	 fmt.Scan(&years)
	 
     futureValue := float64(investmentAmount )* math.Pow( 1 + expectedReturnRate/100 , years)
	 futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Print("Future Value:", futureValue)
	fmt.Println("Future Value(adjusted for Inflation)" , futureRealValue)
}

# 2. Profit-Calculator

package main

import "fmt"

func main(){
	var revenue float64
	var expenses float64
	var taxRate float64 

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("tax Rate: ")
	fmt.Scan(&taxRate )

	ebt := revenue - expenses
	profit := ebt * (1- taxRate/100)
	ratio := ebt/profit

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}

# Future Value with fn basics

package main

import (
	"math"
	"fmt"
)


func main() {
	
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	// fmt.Print("Investment Amount: ")
	outputText("Investment amt: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected return rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)

	// fmt.Print("Future Value:", futureValue)
	// fmt.Println("Future Value(adjusted for Inflation)", futureRealValue)
	// fmt.Printf("Future Value: %f\nFuture Value (adjusted for Inflation): %f", futureValue, futureRealValue)
	// fmt.Printf(`Future Value: %.0f\n
	// Future Value (adjusted for Inflation): %.0f`, futureValue, futureRealValue)
	// fmt.Println("Future Value (adjusted for Inflation): ", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

# Final Future Value with fn

package main

import (
	"math"
	"fmt"
)

const inflationRate = 2.5

func main() {
	
	var investmentAmount float64
	var years float64
	var expectedReturnRate = 5.5

	// fmt.Print("Investment Amount: ")
	outputText("Investment amt: ")
	fmt.Scan(&investmentAmount)

	// fmt.Print("Expected return rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue , futureRealValue := calculateFutureValue(investmentAmount , expectedReturnRate , years)
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Value (adjusted for Inflation): %.1f\n", futureRealValue)

	// fmt.Print("Future Value:", futureValue)
	// fmt.Println("Future Value(adjusted for Inflation)", futureRealValue)
	// fmt.Printf("Future Value: %f\nFuture Value (adjusted for Inflation): %f", futureValue, futureRealValue)
	// fmt.Printf(`Future Value: %.0f\n
	// Future Value (adjusted for Inflation): %.0f`, futureValue, futureRealValue)
	// fmt.Println("Future Value (adjusted for Inflation): ", futureRealValue)

	fmt.Print(formattedFV, formattedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValue(investmentAmount float64 , expectedReturnRate float64 , years float64) (fv float64 , rfv float64){
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}

# if-else & switch 

package main

import "fmt"



func main(){
	var accountBalance  = 1000.0

	fmt.Println("Welcome to Go Bank")

	for {
	fmt.Println("What do you want to do? ")
	fmt.Println("1. Check Balance ")
	fmt.Println("2. Deposit MOney")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

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
	default:
		fmt.Println("Goodbye! ")
		return

	}

	// if choice == 1 {
	// 	fmt.Println("Your balance is : ", accountBalance)
	// } else if choice == 2{
	// 	fmt.Print("Your deposit: ")
	// 	var depositAmount float64
	// 	fmt.Scan(&depositAmount)

	// 	if depositAmount <= 0{
	// 		fmt.Println("Invalid amt . Must be greater than 0.")
	// 		// return
	// 		continue
	// 	}

	// 	accountBalance += depositAmount //accountBalance = accountBalance + depositAMount
	// 	fmt.Println("Balance updated! New Amount: " , accountBalance)
	// } else if choice == 3{
	// 	fmt.Print("Amount you want to withdraw: ")
	// var withdrawmoney float64
	// fmt.Scan(&withdrawmoney)

	// 	if withdrawmoney <= 0{
	// 		fmt.Println("Invalid amt . Must be greater than 0.")
	// 		return
	// 	}

	// 	if withdrawmoney > accountBalance{
	// 		fmt.Println("Invalid amt . cant withdraw more than you have")
	// 		return
	// 	}

	// accountBalance -= withdrawmoney
	// fmt.Println("Balance UPdated ! New amount: ", accountBalance)
	// } else {
	// 	fmt.Println("Goodbye! ")
	// 	break
	// }

	}
	fmt.Println("Thanks for choosing our bank")
}

# Reading & Writing in a File also handling errors

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
	fmt.Println("What do you want to do? ")
	fmt.Println("1. Check Balance ")
	fmt.Println("2. Deposit MOney")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

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
	fmt.Println("Thanks for choosing our bank")
}