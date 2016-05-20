// A purse inspired by an example in "Joe-E: A Security-Oriented Subset of Java"
// http://www.cs.berkeley.edu/~daw/papers/joe-e-ndss10.pdf
package main

import (
	"errors"
	"fmt"
)

type Currency int

type Purse struct {
	c       *Currency
	balance int
}

func NewPurse(c *Currency, balance int) *Purse {
	return &Purse{c, balance}
}

func (p Purse) Balance() int {
	return p.balance
}

func (p Purse) NewEmptyPurse() *Purse {
	return &Purse{c: p.c, balance: 0}
}

func (p *Purse) TakeFrom(src *Purse, amount int) error {
	if p.c != src.c {
		return errors.New("Wrong currency")
	}
	if amount > src.balance {
		return errors.New("Not enough money in purse")
	}
	if amount < 0 || amount+p.balance < 0 {
		return errors.New("Illegal argument")
	}
	src.balance -= amount
	p.balance += amount
	return nil
}

func main() {
	c := new(Currency)
	purse1 := NewPurse(c, 10)
	purse2 := purse1.NewEmptyPurse()

	err := purse2.TakeFrom(purse1, 5)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Balance is", purse1.Balance())

	err = purse2.TakeFrom(purse1, 6)
	if err != nil {
		fmt.Println(err)
	}
}
