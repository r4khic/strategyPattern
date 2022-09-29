package strategyComponents

import (
	"fmt"
	"log"
)

func main() {
	purchase := NewPurchase()
	purchase.RegisterStrategy("PayPal", NewPaypal())
	purchase.RegisterStrategy("CreditCard", NewCreditCard())
	purchase.RegisterStrategy("Bank", NewBank())
	fmt.Println("Enter the payment method you want to use (PayPal, CreditCard or Bank: ")
	var paymentMethod string
	fmt.Scanln(&paymentMethod)

	if err := purchase.Process(paymentMethod); err != nil {
		log.Fatalln(err)
	}
}
