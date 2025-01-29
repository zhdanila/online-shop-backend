package service

import (
	"online-shop-backend/internal/repository"
	"online-shop-backend/internal/service/buyer"
	"online-shop-backend/internal/service/item"
	"online-shop-backend/internal/service/order"
	"online-shop-backend/internal/service/seller"
)

type Service struct {
	Buyer
	Item
	Order
	Seller
}

type Buyer interface {
	CreateBuyer(req *buyer.CreateBuyerRequest) (*buyer.CreateBuyerResponse, error)
	GetBuyer(req *buyer.GetBuyerRequest) (*buyer.GetBuyerResponse, error)
	ListBuyers(req *buyer.ListBuyersRequest) (*buyer.ListBuyersResponse, error)
	UpdateBuyer(req *buyer.UpdateBuyerRequest) (*buyer.UpdateBuyerResponse, error)
	DeleteBuyer(req *buyer.DeleteBuyerRequest) (*buyer.DeleteBuyerResponse, error)
}

type Item interface {
	CreateItem(req *item.CreateItemRequest) (*item.CreateItemResponse, error)
	GetItem(req *item.GetItemRequest) (*item.GetItemResponse, error)
	ListItems(req *item.ListItemsRequest) (*item.ListItemsResponse, error)
	UpdateItem(req *item.UpdateItemRequest) (*item.UpdateItemResponse, error)
	DeleteItem(req *item.DeleteItemRequest) (*item.DeleteItemResponse, error)
}

type Order interface {
	CreateOrder(req *order.CreateOrderRequest) (*order.CreateOrderResponse, error)
	GetOrder(req *order.GetOrderRequest) (*order.GetOrderResponse, error)
	ListOrders(req *order.ListOrderRequest) (*order.ListOrdersResponse, error)
	UpdateOrder(req *order.UpdateOrderRequest) (*order.UpdateOrderResponse, error)
	DeleteOrder(req *order.DeleteOrderRequest) (*order.DeleteOrderResponse, error)
	AddItemToOrder(req *order.AddItemToOrderRequest) (*order.AddItemToOrderResponse, error)
}

type Seller interface {
	CreateSeller(req *seller.CreateSellerRequest) (*seller.CreateSellerResponse, error)
	GetSeller(req *seller.GetSellerRequest) (*seller.GetSellerResponse, error)
	ListSellers(req *seller.ListSellersRequest) (*seller.ListSellersResponse, error)
	UpdateSeller(req *seller.UpdateSellerRequest) (*seller.UpdateSellerResponse, error)
	DeleteSeller(req *seller.DeleteSellerRequest) (*seller.DeleteSellerResponse, error)
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Buyer:  buyer.NewService(repos.Buyer),
		Item:   item.NewService(repos.Item),
		Order:  order.NewService(repos.Order),
		Seller: seller.NewService(repos.Seller),
	}
}
