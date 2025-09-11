package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	//Arrange
	assert := assert.New(t) //assim não precisa passar T em cada assert
	name := "Campaign X"
	content := "Body"
	contacts := []string{"email1@teste.com", "email2@teste.com", "email3@teste.com"}

	// Act
	campaign := NewCampaign(name, content, contacts)

	//Assert
	assert.Equal(campaign.ID, "1")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

/*
testes feitos manualmente
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
*/
