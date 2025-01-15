package ChatUseCaseDTO

type ChatDto struct {
	ConnectionId string `json:"connectionId"`
	Content      string `json:"content"`
	To           string `json:"to"`
}
