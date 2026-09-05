package chapter11

import (
	"fmt"
	"unicode"
)

type AtmState int

const (
	AtmIdle AtmState = iota
	AtmAuthenticated
	AtmDispensing
)

type Atm struct {
	balance  float64
	pin      string
	attempts int
	State    AtmState
}

func NewAtm(balance float64, pin string) *Atm {
	if balance < 0 {
		panic("invalid balance")
	}
	if len(pin) != 4 {
		panic("invalid pin")
	}
	for _, ch := range pin {
		if !unicode.IsDigit(ch) {
			panic("invalid pin")
		}
	}
	return &Atm{balance: balance, pin: pin, State: AtmIdle}
}

func (a *Atm) InsertCard() {
	a.State = AtmIdle
	a.attempts = 0
}

func (a *Atm) EnterPin(pin string) bool {
	if a.State != AtmIdle && a.State != AtmAuthenticated {
		panic("invalid state")
	}
	if pin == a.pin {
		a.State = AtmAuthenticated
		return true
	}
	a.attempts++
	if a.attempts >= 3 {
		a.State = AtmIdle
	}
	return false
}

func (a *Atm) Withdraw(amount float64) float64 {
	if a.State != AtmAuthenticated {
		panic("not authenticated")
	}
	if amount <= 0 || amount > a.balance {
		panic("invalid amount")
	}
	a.State = AtmDispensing
	a.balance -= amount
	a.State = AtmAuthenticated
	return amount
}

func (a *Atm) Balance() float64 {
	if a.State != AtmAuthenticated {
		panic("not authenticated")
	}
	return a.balance
}

func RunQ1106() {
	atm := NewAtm(100.0, "1234")
	atm.InsertCard()
	fmt.Printf("pin ok=%t withdrew=%v balance=%v\n", atm.EnterPin("1234"), atm.Withdraw(40.0), atm.Balance())
}
