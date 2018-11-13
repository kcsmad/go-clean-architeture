package model

type Anime struct {
	Id string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	TotalEpisodes int `json:"totalEpisodes,omitempty"`
}
