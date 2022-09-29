package strategyComponents

import "fmt"

type PayPal struct{}

func NewPaypal() PayPal {
	return PayPal{}
}

func (p PayPal) Pay() error {
	fmt.Println("Processing purchase with PayPal...")
	return nil
}
