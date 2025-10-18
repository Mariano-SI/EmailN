package campaing

import (
	"testing"
)

func TestNewCampaing(t *testing.T) {
	name := "Campain X"
	content := "Body"
	contacts := []string{"email1@email.com", "email1@email.com"}

	campaing := NewCampaing(name, content, contacts)

	if campaing.ID != "1" {
		t.Error("expected 1")
	} else if campaing.Name != name {
		t.Error("expected correct name")
	} else if campaing.Content != content {
		t.Error("expected correct content")
	} else if len(campaing.Contacts) != len(contacts) {
		t.Error("expected correct contacts")
	}
}
