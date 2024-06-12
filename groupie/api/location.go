package api

type LocationResponse struct {
	Index []Location `json:"index"`
}

type Location struct {
	Id        uint64   `json:"id"`
	Locations []string `json:"locations"`
	DatesUrl  string   `json:"dates"`
}
