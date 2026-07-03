// package main

// import "fmt"

// func main(){
// 	var revenue float64
// 	var expenses float64
// 	var taxRate float64 

// 	fmt.Print("Revenue: ")
// 	fmt.Scan(&revenue)

// 	fmt.Print("Expenses: ")
// 	fmt.Scan(&expenses)

// 	fmt.Print("tax Rate: ")
// 	fmt.Scan(&taxRate )

// 	ebt := revenue - expenses
// 	profit := ebt * (1- taxRate/100)
// 	ratio := ebt/profit

// 	fmt.Println(ebt)
// 	fmt.Println(profit)
// 	fmt.Println(ratio)
// }


package main

import "fmt"

func main(){
	// var revenue float64
	// var expenses float64
	// var taxRate float64 

	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)

	// fmt.Print("Expenses: ")
	// fmt.Scan(&expenses)

	// fmt.Print("tax Rate: ")
	// fmt.Scan(&taxRate )

	// ebt := revenue - expenses
	// profit := ebt * (1- taxRate/100)
	// ratio := ebt/profit

	ebt , profit , ratio := CalculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n" , ebt)
	fmt.Printf("%.1f\n" , profit)
	fmt.Printf("%.3f\n" ,ratio)
}

func CalculateFinancials(revenue , expenses , taxRate float64)(float64 , float64 , float64){
	ebt := revenue - expenses
	profit := ebt * (1- taxRate/100)
	ratio := ebt/profit	
	return ebt,profit,ratio
}

func getUserInput(infoText string) float64{
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}