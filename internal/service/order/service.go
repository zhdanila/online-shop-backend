package order

import (
	"errors"
	"online-shop-backend/internal/domain"
	"online-shop-backend/internal/repository"
)

type Service struct {
	repo repository.Order
}

func (s *Service) AddItemToOrder(req *AddItemToOrderRequest) (*AddItemToOrderResponse, error) {
	if req.Quantity < 1 {
		return nil, errors.New("quantity must be at least 1")
	}

	if err := s.repo.AddItemToOrder(domain.OrderItems{
		OrderID:  req.OrderID,
		ItemID:   req.ItemID,
		Quantity: req.Quantity,
		Price:    req.Price,
	}); err != nil {
		return nil, err
	}
	return &AddItemToOrderResponse{}, nil
}

func (s *Service) CreateOrder(req *CreateOrderRequest) (*CreateOrderResponse, error) {
	if err := s.repo.CreateOrder(domain.Order{
		BuyerID:    req.BuyerID,
		TotalPrice: req.TotalPrice,
	}); err != nil {
		return nil, err
	}
	return &CreateOrderResponse{}, nil
}

func (s *Service) GetOrder(req *GetOrderRequest) (*GetOrderResponse, error) {
	order, err := s.repo.GetOrder(req.Id)
	if err != nil {
		return nil, err
	}
	resp := &GetOrderResponse{Order: Order{
		ID:         order.ID,
		BuyerID:    order.BuyerID,
		TotalPrice: order.TotalPrice,
		CreatedAt:  order.CreatedAt,
	}}

	return resp, nil
}

func (s *Service) ListOrders(req *ListOrderRequest) (*ListOrdersResponse, error) {
	res, err := s.repo.ListOrders()
	if err != nil {
		return nil, err
	}

	resp := &ListOrdersResponse{
		Orders: make([]Order, len(res)),
	}
	for _, item := range res {
		resp.Orders = append(resp.Orders, Order{
			ID:         item.ID,
			BuyerID:    item.BuyerID,
			TotalPrice: item.TotalPrice,
			CreatedAt:  item.CreatedAt,
		})
	}

	return resp, nil
}

func (s *Service) UpdateOrder(req *UpdateOrderRequest) (*UpdateOrderResponse, error) {
	if err := s.repo.UpdateOrder(req.Id, domain.Order{
		BuyerID:    req.BuyerID,
		TotalPrice: req.TotalPrice,
	}); err != nil {
		return nil, err
	}

	return &UpdateOrderResponse{}, nil
}

func (s *Service) DeleteOrder(req *DeleteOrderRequest) (*DeleteOrderResponse, error) {
	if err := s.repo.DeleteOrder(req.Id); err != nil {
		return nil, err
	}

	return &DeleteOrderResponse{}, nil
}

func NewService(repo repository.Order) *Service {
	return &Service{repo: repo}
}
