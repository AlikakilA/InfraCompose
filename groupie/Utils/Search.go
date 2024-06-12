package Utils

import (
	"InfraCompose/groupie/api"
	"InfraCompose/groupie/internals"
	"net/http"
	"strconv"
	"strings"
)

func SearchBar(w http.ResponseWriter, r *http.Request) []api.Artist {
	if r.Method != "POST" {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		return nil
	}

	text := strings.ToLower(r.FormValue("searchTerm"))
	artists := internals.GetArtistAPI(w)
	locations := internals.GetLocationAPI(w)
	result := make([]api.Artist, 0)

	for _, artist := range artists {
		artistName := strings.ToLower(artist.Name)
		artistFound := strings.HasPrefix(artistName, text)
		memberFound := false
		albumFound := strings.HasPrefix(strings.ToLower(artist.FirstAlbum), text)
		creationFound := strconv.Itoa(int(artist.CreationDate)) == text

		for _, member := range artist.Members {
			memberName := strings.ToLower(member)
			if strings.HasPrefix(memberName, text) {
				memberFound = true
				break
			}
		}

		if artistFound || memberFound || albumFound || creationFound {
			result = append(result, artist)
			if artistFound && strings.EqualFold(artistName, text) {
				return result
			}
		}
	}

	for _, location := range locations {
		for _, loc := range location.Locations {
			if strings.HasPrefix(strings.ToLower(loc), text) {
				artist := getArtistByID(artists, location.Id)
				if artist != nil {
					result = append(result, *artist)
				}
				break
			}
		}
	}

	if len(result) == 0 {
		return nil
	}

	return result
}

func getArtistByID(artists []api.Artist, id uint64) *api.Artist {
	for _, artist := range artists {
		if artist.Id == id {
			return &artist
		}
	}
	return nil
}
