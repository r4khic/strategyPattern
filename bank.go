package strategyComponents

import "fmt"

type Bank struct{}

func NewBank() Bank {
	return Bank{}
}

func (b Bank) Pay() error {
	fmt.Println("Processing purchase with Bank...")
	return nil
}
