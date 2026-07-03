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

