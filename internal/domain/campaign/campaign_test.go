package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name      = "Campain X"
	content   = "Test Body"
	createdBy = "test@test.com.br"
	contacts  = []string{"email1@email.com", "email1@email.com"}
	fake      = faker.New()
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
	assert.Equal(campaign.CreatedBy, createdBy)
}

func Test_NewCampaign_IDIsNotNew(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	assert.NotNil(campaign.ID)
}
func Test_NewCampaign_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	assert.Greater(campaign.CreatedOn, now)
}
func Test_NewCampaign_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaign("", content, contacts, createdBy)

	assert.Equal("name is required with min 5", error.Error())
}
func Test_NewCampaign_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaign(fake.Lorem().Text(30), content, contacts, createdBy)

	assert.Equal("name is required with max 24", error.Error())
}
func Test_NewCampaign_MustValidateContentMin(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaign(name, "", contacts, createdBy)

	assert.Equal("content is required with min 5", error.Error())
}
func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaign(name, fake.Lorem().Text(1200), contacts, createdBy)

	assert.Equal("content is required with max 1024", error.Error())
}
func Test_NewCampaign_MustValidateContactsMin(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaign(name, content, []string{}, createdBy)

	assert.Equal("contacts is required with min 1", error.Error())
}
func Test_NewCampaign_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaign(name, content, []string{"email invalid"}, createdBy)

	assert.Equal("email is invalid", error.Error())
}

func Test_NewCampaign_InitialStatusMustBePending(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	assert.Equal(Pending, campaign.Status)
}
func Test_NewCampaign_should_return_an_error_if_createdby_is_not_a_email(t *testing.T) {
	assert := assert.New(t)
	invalidEmail := "test"

	_, err := NewCampaign(name, content, contacts, invalidEmail)

	assert.NotNil(err)
	assert.Equal(err.Error(), "createdby is invalid")
}
