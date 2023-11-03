package account

import (
	"errors"
	"fmt"
	"math"
)

type Name string

type Transaction struct {
	Description string
	Amount      float64
}

type Account struct {
	AccNumber     int
	Name          Name
	balance       float64
	InterestRate  float64
	interestValue float64
	age           float64
	transactions  []Transaction
}

func NewAccount(accNumber int, name Name, initialBalance, interestRate float64) *Account {
	return &Account{
		AccNumber:     accNumber,
		Name:          name,
		balance:       initialBalance,
		InterestRate:  interestRate,
		interestValue: 0.0,
		age:           0.0,
		transactions:  []Transaction{},
	}
}

func (a *Account) Deposit(amount float64) {
	if amount > 0 {
		a.balance += amount
		a.transactions = append(a.transactions, Transaction{
			Description: "Deposit",
			Amount:      amount,
		})
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
	a.transactions = append(a.transactions, Transaction{
		Description: "Withdrawal",
		Amount:      -amount,
	})
	return nil
}

func (a *Account) UpdateAge(years float64) {
	a.age += years
}

func (a *Account) SimpleInterest() {
	interest := (a.balance * a.InterestRate * a.age) / 100
	a.interestValue += interest
	a.balance += interest
}

func (a *Account) CompoundInterest() {
	compoundInterest := a.balance * (math.Pow(1+a.InterestRate/100, a.age) - 1)
	a.interestValue += compoundInterest
	a.balance += compoundInterest
}

func (a *Account) PrintTransactionHistory() {
	fmt.Println("Transaction History:")
	for _, transaction := range a.transactions {
		fmt.Printf("Description: %s\n", transaction.Description)
		fmt.Printf("Amount: $%.2f\n", transaction.Amount)
	}
}

func (a *Account) Summary() {
	fmt.Printf("Account Number: %d\n", a.AccNumber)
	fmt.Printf("Account Holder Name: %s\n", a.Name)
	fmt.Printf("Balance: $%.2f\n", a.balance)
	fmt.Printf("Interest Rate: %.2f%%\n", a.InterestRate)
	fmt.Printf("Interest Value: $%.2f\n", a.interestValue)
	fmt.Printf("Age of Account: %.2f years\n", a.age)
}
