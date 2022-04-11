package domain

type BaseResponse[V any] struct {
	Data V      `json:"data"`
	Msg  string `json:"msg"`
}
