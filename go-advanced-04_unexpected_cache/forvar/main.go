package main

import (
	"fmt"
	"sync"
)

type PaymentParams struct {
	UserID string
	Amount int
	Token  string
}

func ProcessPayment(payment *PaymentParams) {
	fmt.Printf("Processing payment for UserID: %s, "+
		"Amount: %d, Token: %s\n",
		payment.UserID, payment.Amount, payment.Token)
}

func main() {
	payments := []*PaymentParams{
		{UserID: "user1", Amount: 100, Token: "111"},
		{UserID: "user2", Amount: 200, Token: "222"},
		{UserID: "user3", Amount: 300, Token: "333"},
	}

	var wg sync.WaitGroup
	wg.Add(len(payments))

	for i := range payments {
		go func() {
			ProcessPayment(payments[i])
			wg.Done()
		}()
	}

	wg.Wait()
}
