package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}
func (r *repositoryMock) Get() []Campaign {

	return []Campaign{}
}

var (
	newCampaign = contract.NewCampaign{
		Name:    "Test Y",
		Content: "Body Hi",
		Emails:  []string{"test1@email.com"},
	}
	service = Service{}
)

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repo := new(repositoryMock)
	repo.On("Save", mock.Anything).Return(nil)

	service.Repository = repo

	id, error := service.Create(newCampaign)

	assert.Nil(error)
	assert.NotNil(id)

}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	newCampaign.Name = ""
	_, err := service.Create(newCampaign)

	assert.False((errors.Is(internalerrors.ErrInternal, err)))

}
func Test_Create_SaveCampaign(t *testing.T) {
	repo := new(repositoryMock)
	repo.On("Save", mock.MatchedBy(func(c *Campaign) bool {
		return c.Name == "Test Y" && c.Content == "Body Hi" && len(c.Contacts) == 1
	})).Return(nil)

	service := Service{Repository: repo}
	newCampaign := contract.NewCampaign{
		Name:    "Test Y",
		Content: "Body Hi",
		Emails:  []string{"test1@email.com"},
	}

	service.Create(newCampaign)

	repo.AssertExpectations(t)
}
func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repo := new(repositoryMock)
	repo.On("Save", mock.Anything).Return(errors.New("internal server error"))

	service := Service{Repository: repo}
	newCampaign := contract.NewCampaign{
		Name:    "Test Y",
		Content: "Body Hi",
		Emails:  []string{"test1@email.com"},
	}

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(err, internalerrors.ErrInternal))
}
