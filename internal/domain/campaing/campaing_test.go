package campaing

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campain X"
	content  = "Test Body"
	contacts = []string{"email1@email.com", "email1@email.com"}
	fake     = faker.New()
)

func Test_NewCampaing_CreateCampaing(t *testing.T) {
	assert := assert.New(t)

	campaing, _ := NewCampaing(name, content, contacts)

	assert.Equal(campaing.Name, name)
	assert.Equal(campaing.Content, content)
	assert.Equal(len(campaing.Contacts), len(contacts))
}

func Test_NewCampaing_IDIsNotNew(t *testing.T) {
	assert := assert.New(t)

	campaing, _ := NewCampaing(name, content, contacts)

	assert.NotNil(campaing.ID)
}
func Test_NewCampaing_CreatedOnMustBeNow(t *testing.T) {
	assert := assert.New(t)
	now := time.Now().Add(-time.Minute)

	campaing, _ := NewCampaing(name, content, contacts)

	assert.Greater(campaing.CreatedOn, now)
}
func Test_NewCampaing_MustValidateNameMin(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaing("", content, contacts)

	assert.Equal("name is required with min 5", error.Error())
}
func Test_NewCampaing_MustValidateNameMax(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaing(fake.Lorem().Text(30), content, contacts)

	assert.Equal("name is required with max 24", error.Error())
}
func Test_NewCampaing_MustValidateContentMin(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaing(name, "", contacts)

	assert.Equal("content is required with min 5", error.Error())
}
func Test_NewCampaing_MustValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaing(name, fake.Lorem().Text(1200), contacts)

	assert.Equal("content is required with max 1024", error.Error())
}
func Test_NewCampaing_MustValidateContactsMin(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaing(name, content, []string{})

	assert.Equal("contacts is required with min 1", error.Error())
}
func Test_NewCampaing_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaing(name, content, []string{"email invalid"})

	assert.Equal("email is invalid", error.Error())
}
