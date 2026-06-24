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

