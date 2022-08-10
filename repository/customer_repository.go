package repository

import (
	"context"
	"fmt"
	"log"

	"github.com/itsapep/yopei-grpc/client/service"
)

type CustomerRepository interface {
	CheckBalance(yopeiId int32) (float32, error)
	DoPayment(yopeiId int32, amount float32) error
}

type customerRepository struct {
	client service.YopeiPaymentClient
}

// CheckBalance implements CustomerRepository
func (c *customerRepository) CheckBalance(yopeiId int32) (float32, error) {
	balance, err := c.client.CheckBalance(context.Background(), &service.CheckBalanceMessage{
		YopeiId: yopeiId,
	})
	if err != nil {
		log.Fatalln("error when checking balance ...", err)
	}
	fmt.Println(balance)
	return 0, err
}

// DoPayment implements CustomerRepository
func (c *customerRepository) DoPayment(yopeiId int32, amount float32) error {
	payment, err := c.client.DoPayment(context.Background(), &service.PaymentMessage{
		YopeiId: yopeiId,
		Amount:  amount,
	})
	if err != nil {
		log.Fatalln("error when doing payment ...", err)
	}
	fmt.Println(payment)
	return nil
}

func NewCustomerRepository(client service.YopeiPaymentClient) CustomerRepository {
	repo := new(customerRepository)
	repo.client = client
	return repo
}
