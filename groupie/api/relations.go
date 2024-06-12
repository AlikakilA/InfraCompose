package api

type RelationResponse struct {
	Index []Relation `json:"index"`
}

type Relation struct {
	Id             uint64              `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
