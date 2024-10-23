package main

import (
	"errors"
	"fmt"
	"time"
)

func chargeForLineItem[T lineItem](newItem T, oldItems []T, balance float64) ([]T, float64, error) {
	// check if the user has a balance with enough funds to be able to pay for the cost of the newItem
	if balance < newItem.GetCost() {
		return nil, 0, errors.New("insufficient funds")
	}
	item := newItem
	oldItems = append(oldItems, item)

	newBalance := balance - newItem.GetCost()
	// Calculate the user's new balance by subtracting the cost of the new item from their balance.
	//this is your second return value.

	return oldItems, newBalance, nil
}

type lineItem interface {
	GetCost() float64
	GetName() string
}

type subscription struct {
	userEmail string
	startDate time.Time
	interval  string
}

func (s subscription) GetName() string {
	return fmt.Sprintf("%s subscription", s.interval)
}

func (s subscription) GetCost() float64 {
	if s.interval == "monthly" {
		return 25.00
	}
	if s.interval == "yearly" {
		return 250.00
	}
	return 0.0
}

type oneTimeUsagePlan struct {
	userEmail        string
	numEmailsAllowed int
}

func (otup oneTimeUsagePlan) GetName() string {
	return fmt.Sprintf("one time usage plan with %v emails", otup.numEmailsAllowed)
}

func (otup oneTimeUsagePlan) GetCost() float64 {
	const costPerEmail = 0.03
	return float64(otup.numEmailsAllowed) * costPerEmail
}
