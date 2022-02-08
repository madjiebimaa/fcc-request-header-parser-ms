package models

type WhoAmI struct {
	IPAddress string `json:"ip_address"`
	Language  string `json:"language"`
	Software  string `json:"software"`
}
