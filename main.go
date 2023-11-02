package main

import (
	"errors"
	"fmt"
)

type Name string

type Account struct {
	AccNumber     int
	Name          Name
	balance       float64
	InterestRate  float64
	interestValue float64
	age           float64
}

func NewAccount(accNumber int, name Name, initialBalance, interestRate float64) *Account {
	return &Account{
		AccNumber:     accNumber,
		Name:          name,
		balance:       initialBalance,
		InterestRate:  interestRate,
		interestValue: 0.0,
		age:           0.0,
	}
}

func (a *Account) Deposit(amount float64) {
	if amount > 0 {
		a.balance += amount
	}
}

func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return errors.New("Invalid withdrawal amount")
	}
	if amount > a.balance {
		return errors.New("Insufficient balance")
	}
	a.balance -= amount
	return nil
}

func (a *Account) ChangeInterestRate(newRate float64) {
	if newRate >= 0 {
		a.InterestRate = newRate
	}
}

func (a *Account) SimpleInterest() {
	interest := (a.balance * a.InterestRate * a.age) / 100
	a.interestValue += interest
	a.balance += interest
}

func (a *Account) Summary() {
	fmt.Printf("Account Number: %d\n", a.AccNumber)
	fmt.Printf("Account Holder Name: %s\n", a.Name)
	fmt.Printf("Balance: $%.2f\n", a.balance)
	fmt.Printf("Interest Rate: %.2f%%\n", a.InterestRate)
	fmt.Printf("Interest Value: $%.2f\n", a.interestValue)
	fmt.Printf("Age of Account: %.2f years\n", a.age)
}
