package buyer

import (
	"online-shop-backend/internal/domain"
	"online-shop-backend/internal/repository"
)

type Service struct {
	repo repository.Buyer
}

func (s Service) CreateBuyer(req *CreateBuyerRequest) (*CreateBuyerResponse, error) {
	if err := s.repo.CreateBuyer(domain.Buyer{
		Name:  req.Name,
		Phone: req.Phone,
	}); err != nil {
		return nil, err
	}

	return &CreateBuyerResponse{}, nil
}

func (s Service) GetBuyer(req *GetBuyerRequest) (*GetBuyerResponse, error) {
	res, err := s.repo.GetBuyer(req.Id)
	if err != nil {
		return nil, err
	}

	resp := &GetBuyerResponse{
		Buyer: Buyer{
			ID:        res.ID,
			Name:      res.Name,
			Phone:     res.Phone,
			CreatedAt: res.CreatedAt,
		},
	}

	return resp, nil
}

func (s Service) UpdateBuyer(req *UpdateBuyerRequest) (*UpdateBuyerResponse, error) {
	if err := s.repo.UpdateBuyer(req.Id, domain.Buyer{
		Name:  req.Name,
		Phone: req.Phone,
	}); err != nil {
		return nil, err
	}

	return &UpdateBuyerResponse{}, nil
}

func (s Service) DeleteBuyer(req *DeleteBuyerRequest) (*DeleteBuyerResponse, error) {
	if err := s.repo.DeleteBuyer(req.Id); err != nil {
		return nil, err
	}

	return &DeleteBuyerResponse{}, nil
}

func NewService(repo repository.Buyer) *Service {
	return &Service{repo: repo}
}
