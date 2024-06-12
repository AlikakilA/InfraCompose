package main

import (
	"InfraCompose/groupie/Utils"
	"InfraCompose/groupie/internals"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		artists := internals.GetArtistAPI(w)
		internals.IndexTemplate(w, r, artists)
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		result := Utils.SearchBar(w, r)
		if result != nil {
			internals.IndexTemplate(w, r, result)
		} else {
			http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		}
	})

	http.HandleFunc("/filters", func(w http.ResponseWriter, r *http.Request) {
		result := Utils.FilterArtists(w, r)
		internals.IndexTemplate(w, r, result)
	})
	//http.HandleFunc("/filters", internals.FiltersHandler)

	http.HandleFunc("/singleArtist", func(w http.ResponseWriter, r *http.Request) {
		artist := internals.ReturnSingleArticle(w, r)
		internals.SingleArtistTemplate(w, artist)
	})

	http.ListenAndServe(":8081", nil)
}
