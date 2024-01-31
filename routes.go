package main

import (
	"database/sql"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func GetExpenses(c *fiber.Ctx) error {
	db:= c.Locals("DB").(*sql.DB)

	expenses,err := FetchExpensesFromDB(db)
	if err!= nil{
		return err
	}

	return c.JSON(expenses)	
}

func CreateExpense(c *fiber.Ctx) error {
	db := c.Locals("DB").(*sql.DB)

	var newExp Expense 
	if err:= c.BodyParser(&newExp); err!= nil{
		return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON format")
	}


	CreatedExp,err := CreateNewExpense(db,newExp)
	if err!= nil{
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(CreatedExp)
}




func GetExpense(c *fiber.Ctx) error {
	db := c.Locals("DB").(*sql.DB)

	id := c.Params("id")

	res,err := FetchExpensesByIDFromDB(db,id)
	if err!= nil{
		return fmt.Errorf("Unable to fetch data")
	}

	return c.JSON(res)
}

func DeleteExpense(c *fiber.Ctx) error {
	db := c.Locals("DB").(*sql.DB)

	id := c.Params("id")

	err := DeleteResponseFromDB(db,id)
	if err!= nil{
		return fmt.Errorf("Unable to Delete the Expense %v",err)
	}

	return c.JSON("EXPENSE DELETED!")
	
}

func UpdateExpense(c *fiber.Ctx) error {
    db := c.Locals("DB").(*sql.DB)
    id := c.Params("id")

    // Parse the updated expense from the request body
    var updatedExp Expense
    if err := c.BodyParser(&updatedExp); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString("Invalid JSON format")
    }

    res, err := UpdateExpenseInDB(db, id, updatedExp)
    if err != nil {
        return fmt.Errorf("Unable to update the expense: %v", err)
    }

    return c.JSON(res)
}