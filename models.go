package main

import "time"

type Expense struct {
	ID uint `json:"id"` 
	Amount int `json:"amount"` 
	Category string `json:"category"` 
	Date time.Time `json:"time"` 
	Description string `json:"description"` 
} 