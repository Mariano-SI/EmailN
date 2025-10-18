package campaing

import (
	"emailn/internal/contract"
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

func Test_Create_Campaing(t *testing.T) {
	assert := assert.New(t)
	service := Service{}
	newCampaing := contract.NewCampaing{
		Name:    "Test Y",
		Content: "Body",
		Emails:  []string{"test1@email.com"},
	}

	id, error := service.Create(newCampaing)

	assert.Nil(error)
	assert.NotNil(id)

}
func Test_Create_SaveCampaing(t *testing.T) {
	newCampaing := contract.NewCampaing{
		Name:    "Test Y",
		Content: "Body",
		Emails:  []string{"test1@email.com"},
	}
	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaing *Campaing) bool {

		if campaing.Name != newCampaing.Name {
			return false
		} else if campaing.Content != newCampaing.Content {
			return false
		} else if len(campaing.Contacts) != len(newCampaing.Emails) {
			return false
		}
		return true
	})).Return(nil)
	service := Service{
		Repository: repositoryMock,
	}

	service.Create(newCampaing)

	repositoryMock.AssertExpectations(t)
}
