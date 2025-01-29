package item

import (
	"online-shop-backend/internal/domain"
	"online-shop-backend/internal/repository"
)

type Service struct {
	repo repository.Item
}

func (s Service) CreateItem(req *CreateItemRequest) (*CreateItemResponse, error) {
	if err := s.repo.CreateItem(domain.Item{
		SellerID:    req.SellerID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}); err != nil {
		return nil, err
	}

	return &CreateItemResponse{}, nil
}

func (s Service) GetItem(req *GetItemRequest) (*GetItemResponse, error) {
	res, err := s.repo.GetItem(req.Id)
	if err != nil {
		return nil, err
	}

	resp := &GetItemResponse{Item: Item{
		ID:          res.ID,
		SellerID:    res.SellerID,
		Name:        res.Name,
		Description: res.Description,
		Price:       res.Price,
		CreatedAt:   res.CreatedAt,
	}}

	return resp, nil
}

func (s Service) ListItems(req *ListItemsRequest) (*ListItemsResponse, error) {
	res, err := s.repo.ListItems()
	if err != nil {
		return nil, err
	}

	resp := &ListItemsResponse{
		Items: make([]Item, len(res)),
	}
	for _, item := range res {
		resp.Items = append(resp.Items, Item{
			ID:          item.ID,
			SellerID:    item.SellerID,
			Name:        item.Name,
			Description: item.Description,
			Price:       item.Price,
			CreatedAt:   item.CreatedAt,
		})
	}

	return resp, nil
}

func (s Service) UpdateItem(req *UpdateItemRequest) (*UpdateItemResponse, error) {
	if err := s.repo.UpdateItem(req.Id, domain.Item{
		SellerID:    req.SellerID,
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
	}); err != nil {
		return nil, err
	}

	return &UpdateItemResponse{}, nil
}

func (s Service) DeleteItem(req *DeleteItemRequest) (*DeleteItemResponse, error) {
	if err := s.repo.DeleteItem(req.Id); err != nil {
		return nil, err
	}

	return &DeleteItemResponse{}, nil
}

func NewService(repo repository.Item) *Service {
	return &Service{repo: repo}
}
