package models

type BOT struct {
	ID        int    `json:"id"`
	IP        string `json:"ip"`
	Consec404 int    `json:"consec404"`
}

type BOTS []BOT
