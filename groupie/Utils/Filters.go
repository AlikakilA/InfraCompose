package Utils

import (
	"InfraCompose/groupie/api"
	"InfraCompose/groupie/internals"
	"net/http"
	"strconv"
	"strings"
)

func FilterArtists(w http.ResponseWriter, r *http.Request) []api.Artist {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		return nil
	}

	minCreationYear, _ := strconv.Atoi(r.FormValue("minCreationYear"))
	maxCreationYear, _ := strconv.Atoi(r.FormValue("maxCreationYear"))
	albumCreation, _ := strconv.Atoi(r.FormValue("input"))
	minYear, _ := strconv.Atoi(r.FormValue("minYear"))
	country := r.FormValue("country")

	artists := internals.GetArtistAPI(w)
	locations := internals.GetLocationAPI(w)
	result := make([]api.Artist, 0)
	for _, artist := range artists {
		if uint16(minCreationYear) <= artist.CreationDate && artist.CreationDate <= uint16(maxCreationYear) {
			result = append(result, artist)
		}
	}

	numMembersValues := r.Form["numMembers[]"]
	var numMembers []int
	for _, value := range numMembersValues {
		numMember, err := strconv.Atoi(value)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusPermanentRedirect)
			continue
		}
		numMembers = append(numMembers, numMember)
	}

	for _, artist := range artists {
		albumYear := extractYearFromAlbum(artist.FirstAlbum)
		if albumYear >= albumCreation && albumYear >= minYear {
			result = append(result, artist)
		}
	}

	if len(numMembers) > 0 {
		for _, artist := range artists {
			if contains(numMembers, len(artist.Members)) {
				result = append(result, artist)
			}
		}
	}

	for _, location := range locations {
		for _, loc := range location.Locations {
			locFound := extractLocation(loc)
			if strings.EqualFold(locFound, country) {
				artist := getArtistByID(artists, location.Id)
				if artist != nil {
					result = append(result, *artist)
				}
				break
			}
		}
	}

	return result
}

func extractYearFromAlbum(album string) int {
	parts := strings.Split(album, "-")
	if len(parts) > 0 {
		year, _ := strconv.Atoi(parts[2])
		return year
	}
	return 0
}

func extractLocation(locations string) string {
	parts := strings.Split(locations, "-")
	if len(parts) > 1 {
		return parts[1]
	}
	return ""
}

func contains(arr []int, num int) bool {
	for _, a := range arr {
		if a == num {
			return true
		}
	}
	return false
}
