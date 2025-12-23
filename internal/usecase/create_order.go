package usecase

import "github.com/devfullcycle/20-CleanArch/internal/entity"

type CreateOrderInputDTO struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type CreateOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type CreateOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewCreateOrderUseCase(orderRepository entity.OrderRepositoryInterface) *CreateOrderUseCase {
	return &CreateOrderUseCase{
		OrderRepository: orderRepository,
	}
}

func (c *CreateOrderUseCase) Execute(input CreateOrderInputDTO) (CreateOrderOutputDTO, error) {
	order, err := entity.NewOrder(input.ID, input.Price, input.Tax)
	if err != nil {
		return CreateOrderOutputDTO{}, err
	}
	err = order.CalculateFinalPrice()
	if err != nil {
		return CreateOrderOutputDTO{}, err
	}
	if err := c.OrderRepository.Save(order); err != nil {
		return CreateOrderOutputDTO{}, err
	}
	dto := CreateOrderOutputDTO{
		ID:         order.ID,
		Price:      order.Price,
		Tax:        order.Tax,
		FinalPrice: order.FinalPrice,
	}
	return dto, nil
}
