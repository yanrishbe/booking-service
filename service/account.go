package service

type Account struct {
	db AccountRepository
}

func NewAccount(repository AccountRepository) *Account {
	return &Account{db: repository}
}
