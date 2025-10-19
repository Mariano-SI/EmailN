package campaing

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

func (r *repositoryMock) Save(campaing *Campaing) error {
	args := r.Called(campaing)
	return args.Error(0)
}

var (
	newCampaing = contract.NewCampaing{
		Name:    "Test Y",
		Content: "Body Hi",
		Emails:  []string{"test1@email.com"},
	}
	service = Service{}
)

func Test_Create_Campaing(t *testing.T) {
	assert := assert.New(t)
	repo := new(repositoryMock)
	repo.On("Save", mock.Anything).Return(nil)

	service.Repository = repo

	id, error := service.Create(newCampaing)

	assert.Nil(error)
	assert.NotNil(id)

}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	newCampaing.Name = ""
	_, err := service.Create(newCampaing)

	assert.False((errors.Is(internalerrors.ErrInternal, err)))

}
func Test_Create_SaveCampaing(t *testing.T) {
	repo := new(repositoryMock)
	repo.On("Save", mock.MatchedBy(func(c *Campaing) bool {
		return c.Name == "Test Y" && c.Content == "Body Hi" && len(c.Contacts) == 1
	})).Return(nil)

	service := Service{Repository: repo}
	newCampaing := contract.NewCampaing{
		Name:    "Test Y",
		Content: "Body Hi",
		Emails:  []string{"test1@email.com"},
	}

	service.Create(newCampaing)

	repo.AssertExpectations(t)
}
func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repo := new(repositoryMock)
	repo.On("Save", mock.Anything).Return(errors.New("internal server error"))

	service := Service{Repository: repo}
	newCampaing := contract.NewCampaing{
		Name:    "Test Y",
		Content: "Body Hi",
		Emails:  []string{"test1@email.com"},
	}

	_, err := service.Create(newCampaing)

	assert.True(errors.Is(err, internalerrors.ErrInternal))
}
