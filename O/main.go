package main

//
//import "fmt"
//
//type PaymentMethod interface {
//	Pay()
//}
//
//type Payment struct{}
//
//func (p Payment) Process(pm PaymentMethod) {
//	pm.Pay()
//}
//
//type CreditCard struct {
//	amount float64
//}
//
//func (cc CreditCard) Pay() {
//	fmt.Printf("Paid %.2f using CreditCard\n", cc.amount)
//}
//
//type KaspiBank struct {
//	amount float64
//	name   string
//}
//
//func (ks KaspiBank) Pay() {
//	a := 2
//	b := 3
//	c := a + b
//	fmt.Println("Hey you, Here is the interface", c)
//	fmt.Printf("Paid %.2f, using %s", ks.amount, ks.name)
//}
//
//func cmd() {
//	p := Payment{}
//	cc := CreditCard{12.23}
//
//	ks := KaspiBank{5.32, "Kaspi bank"}
//
//	p.Process(cc)
//	p.Process(ks)
//
//}
