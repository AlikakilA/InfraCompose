package api

type DateResponse struct {
	Index []Date `json:"index"`
}

type Date struct {
	Id    uint64   `json:"id"`
	Dates []string `json:"dates"`
}
