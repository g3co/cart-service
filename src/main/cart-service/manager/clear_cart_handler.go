package manager

import (
	"main/cart-service/structs"
)

func (m *Manager) ClearCartHandler(client structs.Client, data []byte) (res structs.IMessage, err error) {
	resp := new(structs.AddCartResponse)

	err = m.Repo.ClearCart(client.UserId)
	if err != nil {
		return
	}

	resp.Status = true
	res = resp
	return
}
