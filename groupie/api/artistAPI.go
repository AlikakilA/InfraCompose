package api

type Artist struct {
	Id              uint64   `json:"id"`
	Image           string   `json:"image"`
	Name            string   `json:"name"`
	Members         []string `json:"members"`
	CreationDate    uint16   `json:"creationDate"`
	FirstAlbum      string   `json:"firstAlbum"`
	LocationsUrl    string   `json:"locations"`
	ConcertDatesUrl string   `json:"concertDates"`
	RelationsUrl    string   `json:"relations"`
}
