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


// Program Starts
//       │
//       ▼
// Import Packages
//       │
//       ▼
// Create Constants
//       │
//       ▼
// Create Variables
//       │
//       ▼
// Take User Input
//       │
//       ▼
// Calculate Future Value
//       │
//       ▼
// Calculate Inflation Adjusted Value
//       │
//       ▼
// Print Results
//       │
//       ▼
// Program Ends
