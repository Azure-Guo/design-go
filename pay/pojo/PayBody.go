package pojo

type PayBody struct {
	Account string `json:"account"`
	Type    int    `json:"type"`
	Product string `json:"product"`
	Amount  int    `json:"amount"`
}
