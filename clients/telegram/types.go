package telegram

type UpdateResponse struct {
	OK     bool     `json:"ok"`
	Result []Update `json:"result"`
}

type Update struct {
	ID      int              `json:"update_id"`
	Message *IncomingMessage `json:"message"`
}

type From struct {
	Username string `json:"username"`
}

type Chat struct {
	ID int `json:"id"`
}

type IncomingMessage struct {
	Audio []byte `json:"audio"`
	From  From   `json:"from"`
	Chat  Chat   `json:"chat"`
}
