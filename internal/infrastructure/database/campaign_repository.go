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

func (cr *CampaignRepository) Update(campaign *campaign.Campaign) error {
	tx := cr.Db.Save(campaign)
	return tx.Error
}

func (cr *CampaignRepository) Get() ([]campaign.Campaign, error) {

	var campaings []campaign.Campaign
	tx := cr.Db.Find(&campaings)
	return campaings, tx.Error
}

func (cr *CampaignRepository) GetBy(id string) (*campaign.Campaign, error) {
	var campaign campaign.Campaign

	tx := cr.Db.Preload("Contacts").First(&campaign, "id = ?", id)

	return &campaign, tx.Error
}

func (cr *CampaignRepository) Delete(campaign *campaign.Campaign) error {

	for _, contact := range campaign.Contacts {
		cr.Db.Delete(contact)
	}

	tx := cr.Db.Delete(campaign)
	return tx.Error
}
