package repository

import (
	"GoProjectArche/internal/model"
	"database/sql"
)

/*
communicate with DB
*/
type AccountRepository struct {
	DB *sql.DB
}

func (r *AccountRepository) CreateAccount(acc model.Account) error {
	//query
	query := "INSERT INTO accounts(name, balance) VALUES (?, ?)"
	//excuting query if err return err
	_, err := r.DB.Exec(query, acc.Name, acc.Balance)
	return err
}

func (r *AccountRepository) GetAccountByID(id int) (model.Account, error) {
	var acc model.Account
	query := "SELECT id, name, balance FROM accounts WHERE id = ?"
	//QueryRow() return one row data
	err := r.DB.QueryRow(query, id).Scan(&acc.ID, &acc.Name, &acc.Balance)
	return acc, err
}

/*
note:
r.DB.Excu(): no return only excute the query
r.DB.QueryRow(): returns single row data
r.DB.Query() : returns multi row data
*/
