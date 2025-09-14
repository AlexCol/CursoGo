package campaign

import (
	"emailIn/src/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)
}

func Test_Create_Campaign(t *testing.T) {
	//Arrange
	assert := assert.New(t)
	newCampaign := contract.NewCampaignDto{
		Name:    "Teste Y",
		Content: "Body",
		Emails:  []string{"teste@teste.com"},
	}

	repositoryMock := new(repositoryMock)
	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		if campaign.Name != newCampaign.Name {
			return false
		}
		if campaign.Content != newCampaign.Content {
			return false
		}
		if len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		}

		return true
	})).Return(nil)

	service := Service{Repository: repositoryMock}

	// Act
	id, err := service.Create(newCampaign)

	//Assert
	assert.NotEmpty(id)
	assert.Nil(err)
	repositoryMock.AssertExpectations(t)
}
