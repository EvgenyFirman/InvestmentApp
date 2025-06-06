package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount float64
	years := 10.0
	expectedReturnRate := 5.5

	fmt.Print("Investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Years of investment: ")
	fmt.Scan(&years)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, float64(years))

	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}
