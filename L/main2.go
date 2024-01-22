package main

type Bird interface {
	MakeSound() string
}

type FlyingBird interface {
	Bird
	Fly() string
}

type Pigeon struct{}

func (p *Pigeon) MakeSound() string {
	return "Coo"
}

func (p *Pigeon) Fly() string {
	return "Pigeon is flying."
}

type Penguin struct{}

func (p *Penguin) MakeSound() string {
	return "Squawk"
}
