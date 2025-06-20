package campaign

import (
	"time"

	"github.com/rs/xid"
)

// Definição do tipo Contact
type Contact struct {
	Email string `validade:"email"`
}

// Definição do tipo Campaign
type Campaign struct {

	//Validação de Struct para validar requisiçoes.
	ID       string    `validate:"required"`
	Name     string    `validade:"min=5,max=24"`
	CreateOn time.Time `validate:"required"`
	Content  string    `validate:"min=5,max=1024"`
	Contact  []Contact `validate:"min=1,dive"`
}

// Função para criar uma nova campanha
func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contact := make([]Contact, len(emails))
	for index, emails := range emails {
		contact[index].Email = emails
	}

	return &Campaign{
		ID:       xid.New().String(),
		Name:     name,
		Content:  content,
		CreateOn: time.Now(),
		Contact:  contact,
	}, nil
}
