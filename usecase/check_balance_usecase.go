package usecase

import "github.com/itsapep/yopei-grpc/client/repository"

type CheckBalanceUsecase interface {
	GetBalance(yopeiId int32) (float32, error)
}

type checkBalanceUsecase struct {
	repo repository.CustomerRepository
}

// GetBalance implements CheckBalanceUsecase
func (c *checkBalanceUsecase) GetBalance(yopeiId int32) (float32, error) {
	return c.repo.CheckBalance(yopeiId)
}

func NewCheckBalanceUsecase(repo repository.CustomerRepository) CheckBalanceUsecase {
	return &checkBalanceUsecase{
		repo: repo,
	}
}
