package main

import "fmt"

type PaymentMethod interface {
	Pay()
}

type Payment struct {
	paymentMethod PaymentMethod
}

func (p Payment) Process() {
	p.paymentMethod.Pay()
}

type CreditCard struct {
	amount float64
}

func (cc CreditCard) Pay() {
	fmt.Printf("Paid %.2f using CreditCard\n", cc.amount)
}

type KaspiBank struct {
	amount float64
	name   string
}

func (ks KaspiBank) Pay() {
	a := 2
	b := 3
	c := a + b
	fmt.Println("Hey you, Here is the interface", c)
	fmt.Printf("Paid %.2f, using %s", ks.amount, ks.name)
}

func main() {
	cc := CreditCard{12.23}
	ks := KaspiBank{5.32, "Kaspi bank"}
	paymentMethodKaspi := &Payment{paymentMethod: ks}
	paymentMethodCard := &Payment{paymentMethod: cc}

	paymentMethodCard.Process()
	paymentMethodKaspi.Process()
}
