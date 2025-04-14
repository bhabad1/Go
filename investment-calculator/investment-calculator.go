package main

import (
	"fmt"
	"math"
)
const inflationRate float64 = 6.5
func main(){
	// Type conversion
	// var investmentAmount =1000;
	// var expectedReturnRate = 5.5;
	// var years =10;

	// var futureValue = float64(investmentAmount) * math.Pow( (1+expectedReturnRate/100),float64(years));
	// fmt.Println(futureValue)

	// Explicit Type Convsersion

	// var investmentAmount float64 =1000;
	// var expectedReturnRate = 5.5; //expectedReturnRate :=5.5 (infered)
	// var years float64=10;
	
	// declaring all variables on the same line
	// investmentAmount, years, expectedReturnRate :=10000.0, 10.0, 5.5
	
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64
    printText("Investment Amount: ")
	fmt.Scan(&investmentAmount)
     printText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)
	printText("Years: ")
	fmt.Scan(&years)
	 futureValue := investmentAmount * math.Pow( (1+expectedReturnRate/100),years);
	 futureRealValue:= futureValue / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureValue)
	fmt.Println((futureRealValue))
}

func printText(text string){
	fmt.Printf(text);
}

func futureValues(investmentAmount, expectedReturnRate,years float64) (float64, float64){
	fv:= investmentAmount * math.Pow( (1+expectedReturnRate/100),years);
	frv:= fv / math.Pow(1+inflationRate/100, years)
	return fv, frv ;
}


