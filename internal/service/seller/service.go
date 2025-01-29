package seller

import (
	"online-shop-backend/internal/domain"
	"online-shop-backend/internal/repository"
)

type Service struct {
	repo repository.Seller
}

func (s Service) CreateSeller(req *CreateSellerRequest) (*CreateSellerResponse, error) {
	if err := s.repo.CreateSeller(domain.Seller{
		Name:  req.Name,
		Phone: req.Phone,
	}); err != nil {
		return nil, err
	}

	return &CreateSellerResponse{}, nil
}

func (s Service) GetSeller(req *GetSellerRequest) (*GetSellerResponse, error) {
	res, err := s.repo.GetSeller(req.Id)
	if err != nil {
		return nil, err
	}

	resp := &GetSellerResponse{Seller: Seller{
		ID:        res.ID,
		Name:      res.Name,
		Phone:     res.Phone,
		CreatedAt: res.CreatedAt,
	}}

	return resp, nil
}

func (s Service) ListSellers(req *ListSellersRequest) (*ListSellersResponse, error) {
	res, err := s.repo.ListSellers()
	if err != nil {
		return nil, err
	}

	resp := &ListSellersResponse{
		Sellers: make([]Seller, len(res)),
	}
	for _, seller := range res {
		resp.Sellers = append(resp.Sellers, Seller{
			ID:        seller.ID,
			Name:      seller.Name,
			Phone:     seller.Phone,
			CreatedAt: seller.CreatedAt,
		})
	}

	return resp, nil
}

func (s Service) UpdateSeller(req *UpdateSellerRequest) (*UpdateSellerResponse, error) {
	if err := s.repo.UpdateSeller(req.Id, domain.Seller{
		Name:  req.Name,
		Phone: req.Phone,
	}); err != nil {
		return nil, err
	}

	return &UpdateSellerResponse{}, nil
}

func (s Service) DeleteSeller(req *DeleteSellerRequest) (*DeleteSellerResponse, error) {
	if err := s.repo.DeleteSeller(req.Id); err != nil {
		return nil, err
	}

	return &DeleteSellerResponse{}, nil
}

func NewService(repo repository.Seller) *Service {
	return &Service{repo: repo}
}
