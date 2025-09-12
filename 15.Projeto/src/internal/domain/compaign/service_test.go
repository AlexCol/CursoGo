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
		Content: "",
		Emails:  []string{"teste@teste.com"},
	}

	// Act
	_, err := service.Create(newCampaign)

	//Assert
	assert.Nil(err)
	//assert.NotEmpty(id)
}
