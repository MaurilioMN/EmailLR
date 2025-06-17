package campaign

import (
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
func NewCampaign(name string, content string, emails []string) *Campaign {

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
	}
}
