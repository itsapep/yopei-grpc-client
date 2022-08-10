package manager

import (
	"fmt"
	"log"

	"github.com/itsapep/yopei-grpc/client/config"
	"github.com/itsapep/yopei-grpc/client/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type InfraManager interface {
	YopeiClientConn() service.YopeiPaymentClient
}

type infraManager struct {
	yopeiClient service.YopeiPaymentClient
	cfg         config.Config
}

// YopeiClientConn implements InfraManager
func (i *infraManager) YopeiClientConn() service.YopeiPaymentClient {
	return i.yopeiClient
}

func (i *infraManager) initGRPCClient() {
	dial, err := grpc.Dial(i.cfg.URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("did not connect ...", err)
	}

	client := service.NewYopeiPaymentClient(dial)
	i.yopeiClient = client
	fmt.Println("GRPC client connected ...")
}

func NewInfraManager(config config.Config) InfraManager {
	infra := infraManager{
		cfg: config,
	}
	infra.initGRPCClient()
	return &infra
}
