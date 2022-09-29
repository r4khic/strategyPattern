package strategyComponents

import "fmt"

type Purchase struct {
	paymentMethodStrategies map[string]PaymentMethodStrategy
}

func NewPurchase() Purchase {
	return Purchase{paymentMethodStrategies: make(map[string]PaymentMethodStrategy)}
}

func (p *Purchase) RegisterStrategy(name string, strategy PaymentMethodStrategy) {
	p.paymentMethodStrategies[name] = strategy
}

func (p Purchase) Process(paymentMethod string) error {
	paymentMethodStrategy := p.paymentMethodStrategies[paymentMethod]
	if err := paymentMethodStrategy.Pay(); err != nil {
		return fmt.Errorf("purchase.Process.paymentMethodStrategy.Pay(): %w", err)
	}
	fmt.Printf("Purchase with %s completed\n", paymentMethod)
	return nil
}
