package contract

type NewCampaignDto struct {
	Name    string   `json:"name"`
	Content string   `json:"content"`
	Email   []string `json:"email"`
}