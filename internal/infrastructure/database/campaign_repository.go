package database

import (
	"emailn/internal/domain/campaign"

	"gorm.io/gorm"
)

type CampaignRepository struct {
	Db *gorm.DB
}

func (cr *CampaignRepository) Save(campaign *campaign.Campaign) error {
	tx := cr.Db.Create(campaign)
	return tx.Error
}
func (cr *CampaignRepository) Get() ([]campaign.Campaign, error) {

	var campaings []campaign.Campaign
	tx := cr.Db.Find(&campaings)
	return campaings, tx.Error
}

func (cr *CampaignRepository) GetBy(id string) (*campaign.Campaign, error) {
	var campaign campaign.Campaign

	tx := cr.Db.First(&campaign, "id = ?", id)

	return &campaign, tx.Error
}
