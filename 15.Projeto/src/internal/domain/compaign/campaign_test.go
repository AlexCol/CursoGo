package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Campaign X"
	content  = "Body"
	contacts = []string{"email1@teste.com", "email2@teste.com", "email3@teste.com"}
)

func TestNewCampaign_CreateNewCampaign(t *testing.T) {
	//Arrange
	assert := assert.New(t) //assim não precisa passar T em cada assert
	//!using default values

	// Act
	campaign, _ := NewCampaign(name, content, contacts)

	//Assert
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
}

func TestNewCampaign_IDisNotNil(t *testing.T) {
	//Arrange
	//!using default values

	// Act
	campaign, _ := NewCampaign(name, content, contacts)

	//Assert
	assert.NotNil(t, campaign.ID) //precisa informar t pois não fiz assert := assert.New(t) (o assert a esquerda pode ser qualquer nome)
}

func TestNewCampaign_CreatedOnisNotZero(t *testing.T) {
	//Arrange
	myAssert := assert.New(t)
	//!using default values
	now := time.Now().Add(-time.Minute)

	// Act
	campaign, _ := NewCampaign(name, content, contacts)

	//Assert
	myAssert.False(campaign.CreatedOn.IsZero())
	myAssert.Greater(campaign.CreatedOn, now)
}

func TestNewCampaign_MustValidateName(t *testing.T) {
	//Arrange
	myAssert := assert.New(t)

	// Act
	_, err := NewCampaign("", content, contacts)

	//Assert
	myAssert.Equal("name is required", err.Error())
}

func TestNewCampaign_MustValidateContent(t *testing.T) {
	//Arrange
	myAssert := assert.New(t)

	// Act
	_, err := NewCampaign(content, "", contacts)

	//Assert
	myAssert.Equal("content is required", err.Error())
}

func TestNewCampaign_MustValidateContacts(t *testing.T) {
	//Arrange
	myAssert := assert.New(t)

	// Act
	_, err := NewCampaign(name, content, []string{})

	//Assert
	myAssert.Equal("contact is required", err.Error())
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
