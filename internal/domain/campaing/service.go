package campaing

import (
	"emailn/internal/contract"
)

type Service struct {
	Repository Repository
}

func (s *Service) Create(newCampaing contract.NewCampaing) (string, error) {
	campaing, err := NewCampaing(newCampaing.Name, newCampaing.Content, newCampaing.Emails)
	if err != nil {
		return "", err
	}

	s.Repository.Save(campaing)
	return campaing.ID, nil
}
