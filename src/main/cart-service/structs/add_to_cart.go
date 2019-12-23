package structs

import "encoding/json"

type AddCartRequest struct {
	Items []CartItem `json:"items"`
}

type AddCartResponse struct {
	Status bool `json:"status"`
}

func (res *AddCartResponse) GetJson() (data string) {
	bData, _ := json.Marshal(res)
	data = string(bData)
	return
}
