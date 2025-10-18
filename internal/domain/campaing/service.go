package campaing

import "emailn/internal/contract"

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaing contract.NewCampaing) error {
	return nil
}
