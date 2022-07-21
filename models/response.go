package models

type Response struct {
	Code int `json:"code"`
	Info string `json:"info"`
	Data interface{}  `json:"data"`
}