package utils

import (
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strings"
	"time"
)

// Init the random number generator
func Init() {
	rand.Seed(time.Now().UnixNano())
}

// Choose a random word in a word file
func ChooseRandomWord(gameType string) string {
	var word string
	var numberOfLetters int
	data := make([]byte, 1024)
	switch gameType {
	case "e":
		numberOfLetters = rand.Intn(2) + 4
		if numberOfLetters == 4 {
			file, err := os.Open("hangman/ressources/words/simple4letters.txt")
			if err != nil {
				fmt.Println("Erreur lors de l'ouverture du fichier:", err)
				break
			}
			count, err := file.Read(data)
			if err != nil {
				fmt.Println("Erreur lors de la lecture du fichier:", err)
				break
			}
			words := strings.Fields(string(data[:count]))
			word = words[rand.Intn(len(words))]
		}
		if numberOfLetters == 5 {
			file, err := os.Open("hangman/ressources/words/simple5letters.txt")
			if err != nil {
				fmt.Println("Erreur lors de l'ouverture du fichier:", err)
				break
			}
			count, err := file.Read(data)
			if err != nil {
				fmt.Println("Erreur lors de la lecture du fichier:", err)
				break
			}
			words := strings.Fields(string(data[:count]))
			word = words[rand.Intn(len(words))]
		}
		if numberOfLetters == 6 {
			file, err := os.Open("hangman/ressources/words/simple6letters.txt")
			if err != nil {
				fmt.Println("Erreur lors de l'ouverture du fichier:", err)
				break
			}
			count, err := file.Read(data)
			if err != nil {
				fmt.Println("Erreur lors de la lecture du fichier:", err)
				break
			}
			words := strings.Fields(string(data[:count]))
			word = words[rand.Intn(len(words))]
		}
	case "m":
		numberOfLetters = rand.Intn(1) + 7
		if numberOfLetters == 7 {
			file, err := os.Open("hangman/ressources/words/medium7letters.txt")
			if err != nil {
				fmt.Println("Erreur lors de l'ouverture du fichier:", err)
				break
			}
			count, err := file.Read(data)
			if err != nil {
				fmt.Println("Erreur lors de la lecture du fichier:", err)
				break
			}
			words := strings.Fields(string(data[:count]))
			word = words[rand.Intn(len(words))]
		}
		if numberOfLetters == 8 {
			file, err := os.Open("hangman/ressources/words/medium8letters.txt")
			if err != nil {
				fmt.Println("Erreur lors de l'ouverture du fichier:", err)
				break
			}
			count, err := file.Read(data)
			if err != nil {
				fmt.Println("Erreur lors de la lecture du fichier:", err)
				break
			}
			words := strings.Fields(string(data[:count]))
			word = words[rand.Intn(len(words))]
		}
	case "h":
		numberOfLetters = rand.Intn(5) + 9
		if numberOfLetters >= 9 && numberOfLetters < 15 {
			file, err := os.Open("hangman/ressources/words/hard9_15letters.txt")
			if err != nil {
				fmt.Println("Erreur lors de l'ouverture du fichier:", err)
				break
			}
			count, err := file.Read(data)
			if err != nil {
				fmt.Println("Erreur lors de la lecture du fichier:", err)
				break
			}
			words := strings.Fields(string(data[:count]))
			word = words[rand.Intn(len(words))]
		}
	}
	return word
}
func HideTheWord(wordLenght int, actualword string) string {
	hide := ((wordLenght / 2) - 1)
	var randomInt = make([]int, hide)
	dashes := ""
	for j := 0; j <= hide; j++ {
		for k := 0; k < hide; k++ {
			x := rand.Intn(wordLenght - 1)
			realx := x
			if slices.Contains(randomInt, realx) == false {
				randomInt[k] = realx
			}
		}
	}
	for i := 0; i < wordLenght; i++ {
		if slices.Contains(randomInt, i) == true {
			dashes += string(actualword[i])
		} else {
			dashes += "_"
		}
	}
	return SpaceBefore(dashes)
}

func SpaceBefore(wordWOS string) string {
	wordWS := ""
	for i := 0; i < len(wordWOS); i++ {
		wordWS += string(wordWOS[i]) + " "
	}
	return wordWS
}
func SpaceAfter(wordWS string) string {
	wordWOS := ""
	for i := 0; i < len(wordWS); i++ {
		if string(wordWS[i]) != " " {
			wordWOS += string(wordWS[i])
		} else {
			wordWOS += ""
		}
	}
	return wordWOS
}

func RevealDashes(word string, guess string, dashe string) string {
	var letter string
	letter += guess
	newDashes := ""
	dashes := SpaceAfter(dashe)
	for i := 0; i < len(word); i++ {
		if guess == string(word[i]) {
			newDashes += guess
		} else if string(dashes[i]) == guess {
			newDashes += guess
		} else if string(dashes[i]) != "_" {
			newDashes += string(word[i])
		} else {
			newDashes += "_"
		}
	}
	return SpaceBefore(newDashes)
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
