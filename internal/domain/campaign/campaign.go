package campaign

import (
	internalerrors "emailn/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

const (
	Pending   string = "Pending"
	Started   string = "Started"
	Cancelled string = "Cancelled"
	Deleted   string = "Deleted"
	Done      string = "Done"
)

type Contact struct {
	ID         string `gorm:"size:50"`
	Email      string `validate:"email" gorm:"size:100"`
	CampaignId string `gorm:"size:50"`
}

type Campaign struct {
	ID        string    `validate:"required" gorm:"size:50"`
	Name      string    `validate:"min=5,max=24" gorm:"size:100"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024" gorm:"size:1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string    `gorm:"size:20"`
	CreatedBy string    `validate:"email" gorm:"size:100"`
}

func (c *Campaign) Cancel() {
	c.Status = Cancelled
}
func (c *Campaign) Delete() {
	c.Status = Deleted
}
func (c *Campaign) Done() {
	c.Status = Done
}

func NewCampaign(name string, content string, emails []string, created_by string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index] = Contact{ID: xid.New().String(), Email: email}
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
		Status:    Pending,
		CreatedBy: created_by,
	}

	err := internalerrors.ValidateStruct(campaign)
	if err != nil {
		return nil, err
	}

	return campaign, nil
}
