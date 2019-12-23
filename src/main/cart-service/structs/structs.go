package structs

type Client struct {
	UserId int64 `json:"userId"`
}

type CartItem struct {
	ItemId   int64 `json:"itemId"`
	Quantity int64 `json:"quantity"`
}
