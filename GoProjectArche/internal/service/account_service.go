package service

import (
	"GoProjectArche/internal/model"
	"GoProjectArche/internal/repository"
	"errors"
)

/*
service layer communicate with repo
*/
type AccountService struct {
	Repo *repository.AccountRepository
}

func (s *AccountService) CreateAccount(acc model.Account) error {
	//bussness logics->if fails return
	if acc.Balance < 0 {
		return errors.New("initial balance cannot be negative")
	}
	//sending data to repo
	return s.Repo.CreateAccount(acc)
}

func (s *AccountService) GetAccountByID(id int) (model.Account, error) {
	return s.Repo.GetAccountByID(id)
}
