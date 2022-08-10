package manager

import "github.com/itsapep/yopei-grpc/client/usecase"

type UsecaseManager interface {
	CheckBalanceUsecase() usecase.CheckBalanceUsecase
}

type usecaseManager struct {
	repoManager RepositoryManager
}

// CheckBalanceUsecase implements UsecaseManager
func (u *usecaseManager) CheckBalanceUsecase() usecase.CheckBalanceUsecase {
	return usecase.NewCheckBalanceUsecase(u.repoManager.CustomerRepository())
}

func NewUsecaseManager(repoManager RepositoryManager) UsecaseManager {
	return &usecaseManager{
		repoManager: repoManager,
	}
}
