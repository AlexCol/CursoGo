package campaign

import "testing"

func TestNewCampaign(t *testing.T) {
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@teste.com", "email2@teste.com", "email3@teste.com"}

	campaign := NewCampaign(name, content, contacts)

	if campaign.ID != "1" {
		t.Errorf("Expected %s, got %s", "1", campaign.ID)
	}

	if campaign.Name != name {
		t.Errorf("Wrong name. Expected %s, got %s", name, campaign.Name)
	}

	if campaign.Content != content {
		t.Errorf("Wrong Content. Expected %s, got %s", content, campaign.Content)
	}

	campaignContactsSize := len(campaign.Contacts)
	contactsSize := len(contacts)
	if campaignContactsSize != contactsSize {
		t.Errorf("Contacts list. Expected size %d, got %d", contactsSize, campaignContactsSize)
	}
}
