package campaign

import (
	"emailIn/src/internal/contract"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create_Campaign(t *testing.T) {
	//Arrange
	assert := assert.New(t)
	service := Service{}
	newCampaign := contract.NewCampaignDto{
		Name:    "Teste Y",
		Content: "Body",
		Emails:  []string{"teste@teste.com"},
	}

	// Act
	id, err := service.Create(newCampaign)

	//Assert
	assert.NotEmpty(id)
	assert.Nil(err)
}
