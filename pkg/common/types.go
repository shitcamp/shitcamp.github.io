package common

type Response struct {
	Error string      `json:"error,omitempty"`
	Data  interface{} `json:"data"`
}
