package models

type Robot struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	Status string `json:"status"`
	Battery int `json:"battery"`
}


