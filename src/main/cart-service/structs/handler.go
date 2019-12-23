package structs

type IMessage interface {
	GetJson() string
}

type HandlerFunc func(client Client, body []byte) (res IMessage, err error)
