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
