package types

type ResponseData struct{
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

type ValidKey struct {
	Key string
	KeyType string 
	Action *string
	ActionValue *interface{}
}