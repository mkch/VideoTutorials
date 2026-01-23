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

var pool = sync.Pool{
	New: func() any {
		return &PaymentParams{}
	},
}

func ProcessPayment(payment *PaymentParams) {
	fmt.Printf("Processing payment for UserID: %s, "+
		"Amount: %d, Token: %s\n",
		payment.UserID, payment.Amount, payment.Token)
	pool.Put(payment)
}

func main() {
	for range 19 {
		payments := make(chan *PaymentParams)

		go func() {
			pay1 := pool.Get().(*PaymentParams)
			pay1.UserID = "user1"
			pay1.Amount = 100
			pay1.Token = "111"
			payments <- pay1

			pay2 := pool.Get().(*PaymentParams)
			pay2.UserID = "user2"
			pay2.Amount = 200
			//pay2.Token = "222"
			payments <- pay2

			pay3 := pool.Get().(*PaymentParams)
			pay3.UserID = "user3"
			pay3.Amount = 300
			pay3.Token = "333"
			payments <- pay3

			close(payments)
		}()

		for payment := range payments {
			ProcessPayment(payment)
		}
	}
}
