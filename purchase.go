package strategyComponents

type PaymentMethodStrategy interface {
	Pay() error
}
