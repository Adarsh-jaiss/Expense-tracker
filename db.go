package main

import(
	"database/sql"

	_ "github.com/lib/pq"
	"fmt"
)

func OpenDB() error{
	var err error
	DB,err = sql.Open("postgres","user=postgres dbname=expenses password='Yourpassword' sslmode=disable")
	if err!= nil{
		return err
	}
	fmt.Println("database connected successfully!")
	return nil
}

func CloseDB() error {
	return DB.Close()
}

func CreateTable() error {
	_, err := DB.Exec("CREATE TABLE IF NOT EXISTS EXPENSES(id SERIAL PRIMARY KEY, amount INT NOT NULL, category VARCHAR(255) NOT NULL, time TIMESTAMP NOT NULL, description VARCHAR(255))")
	if err != nil {
		return fmt.Errorf("unable to create table: %v", err)
	}

	fmt.Println("Table created successfully!")
	return nil
}