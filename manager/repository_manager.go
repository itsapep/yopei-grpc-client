package manager

import "github.com/itsapep/yopei-grpc/client/repository"

type RepositoryManager interface {
	CustomerRepository() repository.CustomerRepository
}

type repositoryManager struct {
	infraManager InfraManager
}

// CustomerRepository implements RepositoryManager
func (r *repositoryManager) CustomerRepository() repository.CustomerRepository {
	return repository.NewCustomerRepository(r.infraManager.YopeiClientConn())
}

func NewRepositoryManager(infraManager InfraManager) RepositoryManager {
	return &repositoryManager{
		infraManager: infraManager,
	}
}
