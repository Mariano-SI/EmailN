package database

import (
	"emailn/internal/domain/campaign"
)

type CampaignRepository struct {
	campaigns []campaign.Campaign
}

func (cr *CampaignRepository) Save(campaign *campaign.Campaign) error {
	cr.campaigns = append(cr.campaigns, *campaign)
	return nil
}
func (cr *CampaignRepository) Get() ([]campaign.Campaign, error) {

	return cr.campaigns, nil
}

func (cr *CampaignRepository) GetBy(id string) (*campaign.Campaign, error) {
	var campaign campaign.Campaign

	for _, v := range cr.campaigns {
		if v.ID == id {
			campaign = v
		}
	}

	return &campaign, nil
}
