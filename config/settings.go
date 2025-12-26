package config

type Settings struct {
	Port       int    `json:"port"`
	ApiBaseUrl string `json:"api_base_url"`
	Token      string `json:"token"`
	DbUrl      string `json:"db_url"`
}
