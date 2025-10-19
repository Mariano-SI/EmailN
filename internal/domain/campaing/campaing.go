package campaing

import (
	internalerrors "emailn/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaing struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=24"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"min=5,max=1024"`
	Contacts  []Contact `validate:"min=1,dive"`
}

func NewCampaing(name string, content string, emails []string) (*Campaing, error) {

	contacts := make([]Contact, len(emails))

	for index, email := range emails {
		contacts[index] = Contact{email}
	}

	campaing := &Campaing{
		ID:        xid.New().String(),
		Name:      name,
		Content:   content,
		CreatedOn: time.Now(),
		Contacts:  contacts,
	}

	err := internalerrors.ValidateStruct(campaing)
	if err != nil {
		return nil, err
	}

	return campaing, nil
}
