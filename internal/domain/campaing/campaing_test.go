package campaing

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNewCampaing(t *testing.T) {
	assert := assert.New(t)
	name := "Campain X"
	content := "Body"
	contacts := []string{"email1@email.com", "email1@email.com"}

	campaing := NewCampaing(name, content, contacts)

	assert.Equal(campaing.ID, "1")
	assert.Equal(campaing.Name, name)
	assert.Equal(campaing.Content, content)
	assert.Equal(len(campaing.Contacts), len(contacts))
	assert.Equal(campaing.ID, "1")
}
