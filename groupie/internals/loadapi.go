package internals

import (
	"InfraCompose/groupie/api"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type SingleArtistData struct {
	Artist         api.Artist
	DatesLocations map[string][]string
}

func GetArtistAPI(w http.ResponseWriter) []api.Artist {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		log.Printf("Error fetching data from API: %v", err)
		handleError(w, http.StatusNotFound, "Error fetching data from API")
		return nil
	}
	// Defer signifie que l'on retarde l'éxecution de la fonction "resp.Body.Close()".
	//La fonction "resp.Body.Close()" permet de fermer le "body" (le corps de la réponse HTTP)
	// à la fin de la récupération des données de l'url.
	defer resp.Body.Close()

	var artistsElements []api.Artist
	if err := json.NewDecoder(resp.Body).Decode(&artistsElements); err != nil {
		log.Printf("Error decoding JSON response: %v", err)
		handleError(w, http.StatusNotFound, "Error decoding JSON response")
		return nil
	}
	return artistsElements
}

func GetLocationAPI(w http.ResponseWriter) []api.Location {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		log.Printf("Error fetching data from API: %v", err)
		handleError(w, http.StatusNotFound, "Error fetching data from API")
		return nil
	}

	defer resp.Body.Close()

	var locationsElements api.LocationResponse
	if err := json.NewDecoder(resp.Body).Decode(&locationsElements); err != nil {
		log.Printf("Error decoding JSON response: %v", err)
		handleError(w, http.StatusNotFound, "Error decoding JSON response")
		return nil
	}
	return locationsElements.Index
}

func GetDatesAPI(w http.ResponseWriter) []api.Date {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		log.Printf("Error fetching data from API: %v", err)
		handleError(w, http.StatusNotFound, "Error fetching data from API")
		return nil
	}

	defer resp.Body.Close()

	var dateResponse api.DateResponse
	if err := json.NewDecoder(resp.Body).Decode(&dateResponse); err != nil {
		log.Printf("Error decoding JSON response: %v", err)
		handleError(w, http.StatusNotFound, "Error decoding JSON response")
		return nil
	}
	return dateResponse.Index
}

func getRelationAPI(w http.ResponseWriter) []api.Relation {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		log.Printf("Error fetching data from API: %v", err)
		handleError(w, http.StatusNotFound, "Error fetching data from API")
		return nil
	}

	defer resp.Body.Close()

	var relationElements api.RelationResponse
	if err := json.NewDecoder(resp.Body).Decode(&relationElements); err != nil {
		log.Printf("Error decoding JSON response: %v", err)
		handleError(w, http.StatusNotFound, "Error decoding JSON response")
		return nil
	}
	return relationElements.Index
}

func IndexTemplate(w http.ResponseWriter, r *http.Request, data interface{}) {
	tmplPath := filepath.Join("groupie/template/", "index.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error parsing template")
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error executing template")
		return
	}
}

func SingleArtistTemplate(w http.ResponseWriter, data SingleArtistData) {
	tmplPath := filepath.Join("groupie/template/", "singleArtists.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error parsing template")
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		handleError(w, http.StatusInternalServerError, "Error executing template")
		return
	}
}

func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) SingleArtistData {
	artistsAPI := GetArtistAPI(w)
	relationsAPI := getRelationAPI(w)
	name := r.URL.Query().Get("name")

	for _, artist := range artistsAPI {
		if artist.Name == name {
			for _, relation := range relationsAPI {
				if relation.Id == artist.Id {
					return SingleArtistData{
						Artist:         artist,
						DatesLocations: relation.DatesLocations,
					}
				}
			}
			break
		}
	}
	return SingleArtistData{}
}

func handleError(w http.ResponseWriter, statusCode int, message string) {
	http.Error(w, message, statusCode)
}
