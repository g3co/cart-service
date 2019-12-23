package manager

import (
	"encoding/json"
	"main/cart-service/structs"
)

func (m *Manager) AddToCartHandler(client structs.Client, data []byte) (res structs.IMessage, err error) {
	var rq structs.AddCartRequest

	err = json.Unmarshal(data, &rq)
	if err != nil {
		return
	}

	resp := new(structs.AddCartResponse)

	err = m.Repo.AddToCart(client.UserId, rq.Items)
	if err != nil {
		return
	}

	resp.Status = true
	res = resp
	return
}
