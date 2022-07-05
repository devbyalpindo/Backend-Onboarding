package dto

type ProductDetail struct {
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
	Maker       UserAction `json:"maker"`
	Checker     UserAction `json:"checker"`
	Signer      UserAction `json:"signer"`
}

type UserAction struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
