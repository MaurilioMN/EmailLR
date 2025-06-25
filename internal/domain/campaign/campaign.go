package campaign

import (
	internalerrors "campaing/internal/internalErrors"
	"time"

	"github.com/rs/xid"
)

// Definição do tipo Contact
type Contact struct {
	Email string `validade:"required,email"`
}

// Definição do tipo Campaign
type Campaign struct {

	//Validação de Struct para validar requisiçoes.
	ID       string    `validate:"required"`
	Name     string    `validate:"min=5,max=24"`
	CreateOn time.Time `validate:"required"`
	Content  string    `validate:"min=5,max=1024"`
	Contact  []Contact `validate:"min=1,dive"`
}

// Função para criar uma nova campanha
func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contact := make([]Contact, len(emails))
	for index, email := range emails {
		contact[index].Email = email
	}

	campaign := &Campaign{
		ID:       xid.New().String(),
		Name:     name,
		Content:  content,
		CreateOn: time.Now(),
		Contact:  contact,
	}
	err := internalerrors.ValidationStruct(campaign)
	if err == nil {
		return campaign, nil
	}
	return nil, err
}
