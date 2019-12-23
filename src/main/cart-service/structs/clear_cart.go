package structs

import "encoding/json"

type ClearCartResponse struct {
	Status bool `json:"status"`
}

func (res *ClearCartResponse) GetJson() (data string) {
	bData, _ := json.Marshal(res)
	data = string(bData)
	return
}
