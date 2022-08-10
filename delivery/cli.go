package delivery

import (
	"fmt"

	"github.com/itsapep/yopei-grpc/client/config"
	"github.com/itsapep/yopei-grpc/client/manager"
)

type cli struct {
	usecaseManager manager.UsecaseManager
}

func (c *cli) Run() {
	balance, err := c.usecaseManager.CheckBalanceUsecase().GetBalance(int32(1))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(balance)
}

func CLI() *cli {
	c := config.NewConfig()
	infraManager := manager.NewInfraManager(c)
	repoManager := manager.NewRepositoryManager(infraManager)
	usecaseManager := manager.NewUsecaseManager(repoManager)
	return &cli{
		usecaseManager: usecaseManager,
	}
}
