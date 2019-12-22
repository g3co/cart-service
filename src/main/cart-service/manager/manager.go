package manager

import (
	"log"
	"main/cart-service/structs"
)

type Manager struct {
	Config structs.Config
}

func (m *Manager) Run() {

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
}
