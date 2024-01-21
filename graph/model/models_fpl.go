package model

type FplResponse struct {
	Teams []FplTeam `json:"teams"`
}

type FplTeam struct {
	Code      int    `json:"code"`
	Name      string `json:"name"`
	ShortName string `json:"short_name"`
}
