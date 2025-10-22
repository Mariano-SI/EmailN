package campaign_test

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"testing"

	internalmock "emailn/internal/test/internal-mock"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func newTestCampaign() contract.NewCampaign {
	return contract.NewCampaign{
		Name:    "Test Y",
		Content: "Body Hi",
		Emails:  []string{"test1@email.com"},
	}
}

func Test_Create_Campaign(t *testing.T) {
	assert := assert.New(t)
	repo := new(internalmock.RepositoryMock)
	repo.On("Save", mock.Anything).Return(nil)

	service := campaign.ServiceImp{Repository: repo}
	campaign := newTestCampaign()

	id, err := service.Create(campaign)

	assert.Nil(err)
	assert.NotNil(id)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	repo := new(internalmock.RepositoryMock)

	service := campaign.ServiceImp{Repository: repo}
	campaign := newTestCampaign()
	campaign.Name = ""

	_, err := service.Create(campaign)

	assert.False(errors.Is(internalerrors.ErrInternal, err))
}

func Test_Create_SaveCampaign(t *testing.T) {
	repo := new(internalmock.RepositoryMock)
	repo.On("Save", mock.MatchedBy(func(c *campaign.Campaign) bool {
		return c.Name == "Test Y" && c.Content == "Body Hi" && len(c.Contacts) == 1
	})).Return(nil)

	service := campaign.ServiceImp{Repository: repo}
	campaign := newTestCampaign()

	service.Create(campaign)

	repo.AssertExpectations(t)
}

func Test_Create_ValidateRepositorySave(t *testing.T) {
	assert := assert.New(t)
	repo := new(internalmock.RepositoryMock)
	repo.On("Save", mock.Anything).Return(errors.New("internal server error"))

	service := campaign.ServiceImp{Repository: repo}
	campaign := newTestCampaign()

	_, err := service.Create(campaign)

	assert.True(errors.Is(err, internalerrors.ErrInternal))
}

func Test_GetById_ReturnCampaign(t *testing.T) {
	assert := assert.New(t)
	campaignData := newTestCampaign()
	campaignTest, _ := campaign.NewCampaign(campaignData.Name, campaignData.Content, campaignData.Emails)

	repo := new(internalmock.RepositoryMock)
	repo.On("GetBy", mock.MatchedBy(func(id string) bool {
		return id == campaignTest.ID
	})).Return(campaignTest, nil)

	service := campaign.ServiceImp{Repository: repo}

	campaignReturned, _ := service.GetBy(campaignTest.ID)

	assert.Equal(campaignTest.ID, campaignReturned.ID)
	assert.Equal(campaignTest.Name, campaignReturned.Name)
	assert.Equal(campaignTest.Content, campaignReturned.Content)
	assert.Equal(campaignTest.Status, campaignReturned.Status)
}

func Test_GetById_ReturnErrorWhenSomethingWrongExists(t *testing.T) {
	assert := assert.New(t)
	campaignData := newTestCampaign()
	campaignTest, _ := campaign.NewCampaign(campaignData.Name, campaignData.Content, campaignData.Emails)

	repo := new(internalmock.RepositoryMock)
	repo.On("GetBy", mock.Anything).Return(nil, errors.New("Something Wrong"))

	service := campaign.ServiceImp{Repository: repo}

	_, err := service.GetBy(campaignTest.ID)

	assert.Equal(internalerrors.ErrInternal.Error(), err.Error())
}
