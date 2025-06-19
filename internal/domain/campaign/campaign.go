package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

// Definição do tipo Contact
type Contact struct {
	Email string
}

// Definição do tipo Campaign
type Campaign struct {
	ID       string
	Name     string
	CreateOn time.Time
	Content  string
	Contact  []Contact
}

// Função para criar uma nova campanha
func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	// Validação de dominio No campaign_test.gos
	if name == "" {
		return nil, errors.New("name is required")
	} else if content == "" {
		return nil, errors.New("content is required")
	} else if len(emails) == 0 {
		return nil, errors.New("contact is required")
	}

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
