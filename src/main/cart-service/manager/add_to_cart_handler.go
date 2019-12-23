package manager

import (
	"main/cart-service/structs"
)

func (m *Manager) GetCartHandler(client structs.Client, data []byte) (res structs.IMessage, err error) {
	resp := new(structs.GetCartResponse)

	resp.Items, err = m.Repo.GetCart(client.UserId)
	if err != nil {
		return
	}

	res = resp
	return
}
