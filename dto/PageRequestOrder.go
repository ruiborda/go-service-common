package dto

type PageRequestOrder struct {
	By    string `json:"by" form:"by"`
	Order string `json:"order" form:"order"`
}
