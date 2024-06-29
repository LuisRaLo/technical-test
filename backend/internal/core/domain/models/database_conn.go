package models

type DBConfig struct {
	Write Write `json:"write"`
	Read  Read  `json:"read"`
}

type Write struct {
	Host     string `json:"host"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Port     int    `json:"port"`
}

type Read struct {
	Host     string `json:"host"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Port     int    `json:"port"`
}
