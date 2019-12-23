package structs

import "encoding/json"

type GetCartResponse struct {
	Items []CartItem `json:"items"`
}

func (res *GetCartResponse) GetJson() (data string) {
	bData, _ := json.Marshal(res)
	data = string(bData)
	return
}
