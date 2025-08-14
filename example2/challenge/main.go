package main

import (
	"fmt"
	"sync"
)

type Income struct {
	Source string
	Amount int
}

func getTotalAmountByIncome(income Income, bankBalance *int, wg *sync.WaitGroup, mtx *sync.Mutex) {
	defer wg.Done()
	var totalByIncome int
	//Si se hace el lock y unlock fuera del for
	//se pierde paralelismo ya que al estarse ejecutando las goroutines de forma concurrente
	// se ejecutarian de una en una, en vez de todas a la vez
	for week := 1; week <= 52; week++ {
		totalByIncome += income.Amount
		mtx.Lock()
		*bankBalance += income.Amount
		mtx.Unlock()
		fmt.Printf("[%s] -> On week %d, tou earned $%d from %s \n", income.Source, week, totalByIncome, income.Source)
	}
}
func main() {

	var bankBalance int
	var mtx sync.Mutex
	var wg sync.WaitGroup

	fmt.Println("************************************")
	fmt.Println("Initial account balance")
	fmt.Printf("Current balance : %d \n", bankBalance)
	fmt.Println("************************************")

	incomes := []Income{
		{Source: "Main job", Amount: 500},
		{Source: "Gifts", Amount: 10},
		{Source: "Part time job", Amount: 50},
		{Source: "Investments", Amount: 100},
	}
	wg.Add(len(incomes))

	for _, income := range incomes {
		go getTotalAmountByIncome(income, &bankBalance, &wg, &mtx)
	}
	wg.Wait()
	fmt.Println()
	fmt.Println("******************************************")
	fmt.Println("Final account balance")
	fmt.Printf("You have earned $%d in 52 weeks (1 year)  \n", bankBalance)
	fmt.Println("******************************************")
	fmt.Println()

}
