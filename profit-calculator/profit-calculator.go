package main

import "fmt"

func main(){

	var revenue, expenses, tax_rate float64

	fmt.Print("Total Revenue: ");
	fmt.Scan(&revenue);

	fmt.Print("Total Expenses: ");
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&tax_rate)

	 ebt := revenue-expenses
	//  profit:= ebt -( ebt*tax_rate/100);
	profit := ebt *(1-tax_rate/100)

	 ratio:= ebt/profit

	 fmt.Println("EBT: ",ebt )
	 fmt.Println("Profit: ",profit)
	 fmt.Println("Ratio (EBT/Profit): ",ratio)

}

