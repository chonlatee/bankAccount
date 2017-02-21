package main

import (
	"sync"

	"fmt"

	bank "github.com/chonlatee/bankAccount/bank"
)

func main() {
	var wg sync.WaitGroup

	for index := 0; index < 5; index++ {
		wg.Add(2)
		go func() {
			defer wg.Done()
			fmt.Println("Add 100")
			bank.Balance.Deposit(100)
			fmt.Printf("Balance after deposit 50: %v \n", bank.Balance.Display())
		}()

		go func() {
			defer wg.Done()
			fmt.Println("Add 50")
			bank.Balance.Deposit(50)
			fmt.Printf("Balance after deposit 50: %v \n", bank.Balance.Display())
		}()
	}

	wg.Wait()

}
