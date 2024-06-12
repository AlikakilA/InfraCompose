package main

import (
	"InfraCompose/hangman/utils"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
)

var (
	tab    []string
	mot    string
	word   string
	stage  = len(tab)
	song   string
	faults int
)

type Page struct {
	Title       string
	Content     []string
	LetterInput []string
	Stage       int
	Mot         string
	entre       string
	Image       string
	Level       string
	NHide       string
	Song        string
	LiSong      string
}
type Log struct {
	UserName string `json:"name"`
	Password string `json:"age"`
	Win      int    `json:"win"`
	Mess     string
	Url      string
}

func save(filename string, data *Log) error {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filename, jsonData, 0644)
	if err != nil {
		return err
	}

	return nil
}
func load(filename string) (*Log, error) {
	jsonData, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var data Log
	err = json.Unmarshal(jsonData, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func choice(mot string) string {
	if len(mot) < 7 && len(mot) > 3 {
		return "/Facile"
	} else if len(mot) >= 7 && len(mot) <= 9 {
		return "/Moyen"
	} else if len(mot) < 3 {
		return "/Ez"
	} else {
		return "/Difficile"
	}
}
func ChangeIMG(stage int) string {
	path := ""
	switch stage {
	case 0:
		path = "ressources/images/pics/Hang1.png"
	case 1:
		path = "ressources/images/pics/Hang2.png"
	case 2:
		path = "ressources/images/pics/Hang3.png"
	case 3:
		path = "ressources/images/pics/Hang4.png"
	case 4:
		path = "ressources/images/pics/Hang5.png"
	case 5:
		path = "ressources/images/pics/Hang6.png"
	case 6:
		path = "ressources/images/pics/Hang7.png"
	case 7:
		path = "ressources/images/pics/Hang8.png"
	}
	return path
}
func ChangeMeme(stage int) string {
	path := ""
	nbrErr := 0
	n := rand.Intn(2) + 1
	if stage > nbrErr {
		switch n {
		case 1:
			path = "ressources/song/huh.mp3"
		case 2:
			path = "ressources/song/uuh.mp3"
		}
	}
	nbrErr++
	return path
}

func Start(w http.ResponseWriter, r *http.Request) {
	data := Page{}
	tmpl, _ := template.ParseGlob("hangman/template/index.html")
	err := tmpl.ExecuteTemplate(w, "index.html", data)
	if err != nil {
		return
	}

}

func Choix(w http.ResponseWriter, r *http.Request) {
	data := Page{
		Title:   word,
		Content: tab,
		//LetterInput: letterInput,
		entre: r.FormValue("entre"),
	}
	tmpl, _ := template.ParseGlob("hangman/template/choice.html")
	err := tmpl.ExecuteTemplate(w, "choice.html", data)
	if err != nil {
		return
	}
}

func Game(w http.ResponseWriter, r *http.Request) {
	test := ""
	letterInput := r.FormValue("letterInput")
	tmpl, _ := template.ParseGlob("hangman/template/game.html")
	if strings.Contains(mot, letterInput) == true && word != "BRAVO" {
		updatedDashes := utils.RevealDashes(mot, letterInput, word)
		word = updatedDashes
	} else if strings.Contains(mot, letterInput) == false && len(letterInput) < 2 {
		tab = append(tab, letterInput)
		faults++
	} else if len(letterInput) > 1 && strings.Contains(mot, letterInput) == false {
		tab = append(tab, letterInput)
		faults += 2
	}
	if utils.SpaceAfter(word) == mot {
		word = "BRAVO"
		song = "ressources/song/sardines.mp3"

	} else {
		song = ""
	}
	if faults >= 7 {
		word = "PERDU"
		song = "ressources/song/flute.mp3"
		test = mot
	}
	data := Page{
		Title:   word,
		Content: tab,
		Image:   ChangeIMG(len(tab)),
		Mot:     test,
		Stage:   7 - faults,
		entre:   r.FormValue("entre"),
		Level:   choice(mot),
		Song:    song,
		LiSong:  ChangeMeme(len(tab)),
	}
	err := tmpl.ExecuteTemplate(w, "game.html", data)
	if err != nil {
		return
	}

}

func Facile(w http.ResponseWriter, r *http.Request) {
	mot = utils.ChooseRandomWord("e")
	tab = []string{}
	faults = 0
	word = utils.HideTheWord(len(mot), mot)
	data := Page{
		Title:   word,
		Content: tab,
		Image:   ChangeIMG(len(tab)),
		Stage:   7 - len(tab),
		entre:   r.FormValue("entre"),
		Level:   "/Facile",
		NHide:   mot,
	}
	tmpl, _ := template.ParseGlob("hangman/template/game.html")
	err := tmpl.ExecuteTemplate(w, "game.html", data)
	if err != nil {
		fmt.Println("error")
	}
}
func Moyen(w http.ResponseWriter, r *http.Request) {
	mot = utils.ChooseRandomWord("m")
	tab = []string{}
	faults = 0
	word = utils.HideTheWord(len(mot), mot)
	data := Page{
		Title:   word,
		Content: tab,
		Image:   ChangeIMG(len(tab)),
		Stage:   7 - len(tab),
		entre:   r.FormValue("entre"),
		Level:   "/Moyen",
		NHide:   mot,
	}
	tmpl, _ := template.ParseGlob("hangman/template/game.html")
	err := tmpl.ExecuteTemplate(w, "game.html", data)
	if err != nil {
		fmt.Println("error")
	}
}
func Difficile(w http.ResponseWriter, r *http.Request) {
	faults = 0
	mot = utils.ChooseRandomWord("h")
	tab = []string{}
	word = utils.HideTheWord(len(mot), mot)
	data := Page{
		Title:   word,
		Content: tab,
		Image:   ChangeIMG(len(tab)),
		Stage:   7 - len(tab),
		entre:   r.FormValue("entre"),
		Level:   "/Difficile",
	}
	tmpl, _ := template.ParseGlob("hangman/template/game.html")
	err := tmpl.ExecuteTemplate(w, "game.html", data)
	if err != nil {
		fmt.Println("error")
	}
}
func Login(w http.ResponseWriter, r *http.Request) {
	var (
		name     string
		password string
	)
	data := &Log{
		UserName: r.FormValue(name),
		Password: r.FormValue(password),
		Mess:     "Username",
		Url:      "/Login",
	}
	tmpl, _ := template.ParseGlob("./template/Login.html")
	err := tmpl.ExecuteTemplate(w, "Login.html", data)
	if err != nil {
		fmt.Println("error")
	}
	if r.Method == "post" && r.FormValue(name) == "" {
		if r.FormValue(name) == "" || r.FormValue(password) == "" {
			http.Redirect(w, r, "/Login", 1)
			data = &Log{
				UserName: r.FormValue(name),
				Password: r.FormValue(password),
				Mess:     "remplis ta mère",
				Url:      "/Choix",
			}
		}
	}
	fmt.Println(name, password)
	filename := name
	save(filename, data)
	err = save(filename, data)
	if err != nil {
		fmt.Println("Erreur lors de la sauvegarde des données :", err)
		return
	}
	a, err := load(name)
	if err != nil {
		fmt.Println("Erreur lors du chargement des données :", err)
		return
	}
	fmt.Println("Nome:", a.UserName)
	fmt.Println("Nome:", a.Password)
}

func main() {
	http.HandleFunc("/", Start)
	//http.HandleFunc("/Login", Login)
	http.HandleFunc("/Choix", Choix)
	http.HandleFunc("/start", Game)
	http.HandleFunc("/Facile", Facile)
	http.HandleFunc("/Moyen", Moyen)
	http.HandleFunc("/Difficile", Difficile)
	http.Handle("/ressources/", http.StripPrefix("/ressources/", http.FileServer(http.Dir("hangman/ressources"))))
	http.Handle("/template/", http.StripPrefix("/template/", http.FileServer(http.Dir("hangman/template"))))
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("hangman/static"))))
	http.ListenAndServe(":8080", nil)

}
