package usecase

import (
	"github.com/chrismarsilva/cms.golang.teste.intensivo/internal/order/entity"
)

type GetTotalOutputDTO struct {
	Total int
}

type GetTotalUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewGetTotalUseCase(orderRepository entity.OrderRepositoryInterface) *GetTotalUseCase {
	return &GetTotalUseCase{
		OrderRepository: orderRepository,
	}
}

func (c *GetTotalUseCase) Execute() (*GetTotalOutputDTO, error) {
	total, err := c.OrderRepository.GetTotal()
	if err != nil {
		return nil, err
	}

	output := &GetTotalOutputDTO{Total: total}
	return output, nil
}
