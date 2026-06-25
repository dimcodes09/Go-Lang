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