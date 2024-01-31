package main

import (
	"database/sql"
	"net/http"
	"fmt"
)

func FetchExpensesFromDB(db *sql.DB) ([]Expense,error) {
	query := "SELECT id,amount,category,time,description FROM EXPENSES"
	rows,err := db.Query(query)
	if err!=nil{
		return nil,fmt.Errorf("Unable to fetch expenses from DB %v",http.StatusInternalServerError)
	}
	defer rows.Close()

	expenses := []Expense{}
	for rows.Next() {
		var exp Expense
		if err := rows.Scan(&exp.ID,&exp.Amount,&exp.Category,&exp.Date,&exp.Description); err!= nil{
			return nil,err
		}
		expenses = append(expenses, exp)
	}

	if err := rows.Err();err!= nil{
		return nil,err
	}
	
	return expenses,nil
}

func CreateNewExpense(db *sql.DB, newExp Expense) (Expense, error) {
    query := "INSERT INTO EXPENSES(amount, category, time, description) VALUES($1, $2, $3, $4) RETURNING id, amount, category, time, description"
    
    err := db.QueryRow(query, newExp.Amount, newExp.Category, newExp.Date, newExp.Description).
        Scan(&newExp.Amount, &newExp.Category, &newExp.Date, &newExp.Description)

    if err != nil {
        fmt.Println("Error executing query:", err)
        return Expense{}, fmt.Errorf("Unable to create new expense: %v", err)
    }

    fmt.Println("New expense created successfully:", newExp)

    return newExp, nil
}

func FetchExpensesByIDFromDB(db *sql.DB, id string) (Expense,error) {
	var Exp Expense
	
	query := "SELECT * FROM expenses WHERE ID=$1"

	err := db.QueryRow(query,id).Scan(&Exp.ID,&Exp.Amount,&Exp.Category,&Exp.Date,&Exp.Description)
	if err!= nil{
		if err == sql.ErrNoRows{
			return Expense{},fmt.Errorf("Expense not found")
		}
		return Expense{},fmt.Errorf("Error fetching the Expense %+v",err)
	}
	
	return Exp,nil
	
}

func DeleteResponseFromDB(db *sql.DB, id string) error {
	query := "DELETE FROM expenses WHERE id=$1"

	_, err := db.Exec(query,id)
	if err!= nil{
		if err == sql.ErrNoRows{
			return fmt.Errorf("Expense not found")
		}
		return fmt.Errorf("Error Deleting the Expense  %+v",err)
	}
	return nil
}

func UpdateExpenseInDB(db *sql.DB, id string, updatedExp Expense) (Expense, error) {
    query := "UPDATE expenses SET amount=$1, category=$2, time=$3, description=$4 WHERE id=$5 RETURNING amount, category, time, description"

    res, err := db.Exec(query, updatedExp.Amount, updatedExp.Category, updatedExp.Date, updatedExp.Description, id)

    if err != nil {
        return Expense{}, fmt.Errorf("Error updating the Expense: %v", err)
    }

    rowsAffected, _ := res.RowsAffected()
    if rowsAffected == 0 {
        return Expense{}, fmt.Errorf("Expense not found")
    }

    return updatedExp, nil
}