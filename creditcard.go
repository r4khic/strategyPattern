package strategyComponents

import "fmt"

type CreditCard struct{}

func NewCreditCard() CreditCard {
	return CreditCard{}
}

func (c CreditCard) Pay() error {
	fmt.Println("Processing purchase with CreditCard...")
	return nil
}
