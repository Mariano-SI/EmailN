package campaing

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campain X"
	content  = "Body"
	contacts = []string{"email1@email.com", "email1@email.com"}
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
func Test_NewCampaing_MustValidateName(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaing("", content, contacts)

	assert.Equal("name is required", error.Error())
}
func Test_NewCampaing_MustValidateContent(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaing(name, "", contacts)

	assert.Equal("content is required", error.Error())
}
func Test_NewCampaing_MustValidateContacts(t *testing.T) {
	assert := assert.New(t)

	_, error := NewCampaing(name, content, []string{})

	assert.Equal("contacts is required", error.Error())
}
