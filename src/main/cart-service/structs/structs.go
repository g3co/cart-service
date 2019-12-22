package structs

type TestRequest struct {
	Token string `json:"token"`
	Body  string `json:"body"`
}

type TestRequestBody struct {
	RequestId int64 `json:"requestId"`
}
