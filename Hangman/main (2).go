package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

var fs = http.FileServer(http.Dir("./static/resources"))
var tmpl = template.Must(template.ParseFiles("index.html"))

type Hangman struct {
	word          string
	Guesses       int
	WrongGuesses2 []string
	state         bool
	guess         string
	draw          []string
	show          []string
	Display       string
	Message       string
	Hangman       string
	Win           int
	Ratio         int
}

var win_next = 0
var win_all = 0
var compteur = 0
var Hangman2 string
var Message = ""
var Messages = ""
var show = []string{}
var axel bool = false
var result = ""
var Guesses2 int = 10
var wrongGuesses []string
var goodGuesses2 []string
var word = ""
var guess = ""
var state = true
var Generate_Random_Word int = 0
var final = ""
var word_guess = ""
var victory = 0
var click = 0
var GuessesLeft = 10
var Draw = []string{"=========", "        |\n	|\n	|\n	|\n	|\n=========", "  +-----+\n        |\n	|\n	|\n	|\n	|\n=========", "  +-----+\n  |     |\n	|\n	|\n	|\n	|\n=========", "  +-----+\n  |     |\n  O	|\n	|\n	|\n	|\n=========", "  +-----+\n  |     |\n  O	|\n  |	|\n	|\n	|\n=========", "  +-----+\n  |     |\n  O	|\n /|	|\n	|\n	|\n=========", "  +-----+\n  |     |\n  O	|\n /|\\    |\n	|\n	|\n=========", "  +-----+\n  |     |\n  O	|\n /|\\    |\n /	|\n	|\n=========", "  +-----+\n  |     |\n  O	|\n /|\\    |\n / \\    |\n	|\n=========", "  +-----+\n  |     |\n  O	|\n /|\\    |\n / \\    |\n	|\n=========\nTU ES MORT"}

func restart() {
	compteur = 0
	Hangman2 = ""
	Message = ""
	Messages = ""
	show = []string{}
	axel = false
	result = ""
	Guesses2 = 10
	wrongGuesses = []string{}
	goodGuesses2 = []string{}
	word = ""
	guess = ""
	state = true
	Generate_Random_Word = 0
	final = ""
	word_guess = ""
	victory = 0
	GuessesLeft = 10
}
func Show_Hide(guess, word string) string {
	final = ""
	for i := 0; i < len(guess); i++ {
		if strings.Contains(word, string(guess[i])) {
			final += string(guess[i])
		} else {
			final += " _ "
		}
	}
	return final
}
func random() {
	if click == 0 {
		if Generate_Random_Word == 0 {
			rand.Seed(time.Now().UnixNano())
			data, err := ioutil.ReadFile("word/words.txt")
			if err != nil {
				os.Exit(0)
			}

			words := string(data)
			wordsList := strings.Fields(words)
			guess = wordsList[rand.Intn(len(wordsList))]
			Generate_Random_Word = 1
			state = true
		}
	}
	if click == 2 {
		if Generate_Random_Word == 0 {
			rand.Seed(time.Now().UnixNano())
			data, err := ioutil.ReadFile("word/words2.txt")
			if err != nil {
				os.Exit(0)
			}

			words := string(data)
			wordsList := strings.Fields(words)
			guess = wordsList[rand.Intn(len(wordsList))]
			Generate_Random_Word = 1
			state = true
		}
	}
	if click == 3 {
		if Generate_Random_Word == 0 {
			rand.Seed(time.Now().UnixNano())
			data, err := ioutil.ReadFile("word/words3.txt")
			if err != nil {
				os.Exit(0)
			}

			words := string(data)
			wordsList := strings.Fields(words)
			guess = wordsList[rand.Intn(len(wordsList))]
			Generate_Random_Word = 1
			state = true
		}
	}
}

const (
	HOST = "localhost"
	PORT = ":8888"
)

func css(w http.ResponseWriter, r *http.Request) {
	//if r.Method != http.MethodPost {
	//	tmpl.Execute(w, nil)
	//	return
	//}
	//code du pendue en lui-même
	word = r.FormValue("w")
	axel = false
	//code du pendue en lui-même
	if word == "" {
		random()
	}
	if len(word) == 1 {
		if word[0] >= 'A' && word[0] <= 'Z' {
			result += string(word[0] + 32)
		}
		if word[0] >= 'a' && word[0] <= 'z' {
			result += string(word[0])
		}
		for i := 0; i < len(guess); i++ {
			for x := 0; x < len(goodGuesses2); x++ {
				if result == goodGuesses2[x] && axel == false {
					// fmt.Fprintln(w, "Bien joué ! Mais déja donner dommage")
					Messages = "Bien joué ! Mais déja donner dommage"
					axel = true
				}
			}
			if result[0] == guess[i] && axel == false {
				// fmt.Fprintln(w, "Tu as trouver une lettre")
				Messages = "Tu as trouver une lettre"
				show = append(show, word)
				goodGuesses2 = append(goodGuesses2, result)
				word_guess += result
				axel = true
			} else {
				if axel == false {
					axel = false
				}
			}
		}
	}
	for z := 0; z < len(wrongGuesses); z++ {
		if result == wrongGuesses[z] && axel == false {
			// fmt.Fprintln(w, "Mauvaise Réponse ! Mais tu as déja donner cette lettre, quelle veinard")
			Messages = "Mauvaise Réponse ! Mais tu as déja donner cette lettre, quelle veinard"
			axel = true
		}
	}
	if axel == false && len(word) == 1 {
		Guesses2 -= 1
		show = append(show, word)
		wrongGuesses = append(wrongGuesses, word)
		// fmt.Fprintln(w, "Dommage mauvaise lettre")
		Messages = "Dommage mauvaise lettre"
		// fmt.Fprintln(w, "tu as perdue 1pv il te reste", Guesses, "pv")
	}
	result = ""
	if len(word) >= 2 {
		for n := 0; n < len(word); n++ {
			if word[n] >= 'A' && word[n] <= 'Z' {
				result += string(word[n] + 32)
			}
			if word[n] >= 'a' && word[n] <= 'z' {
				result += string(word[n])
			}
		}
		if result == guess || Show_Hide(guess, word_guess) == guess {
			// fmt.Fprintln(w, "Tu as gagné boloss")
			Messages = "Tu as gagné"
			axel = true
			word_guess += result
			victory = 1
			restart()
		} else {
			for z := 0; z < len(wrongGuesses); z++ {
				if result == wrongGuesses[z] && axel == false {
					// fmt.Fprintln(w, "Mauvaise Réponse ! Mais tu as déja donner cette lettre, quelle veinard")
					Messages = "Mauvaise Réponse ! Mais tu as déja donner cette lettre, quelle veinard"
					axel = true
				}
			}
			if axel == false && len(word) >= 2 {
				result = ""
				Guesses2 -= 2
				wrongGuesses = append(wrongGuesses, word)
				show = append(show, word)
				// fmt.Fprintln(w, "Dommage mauvaise mot")
				// fmt.Fprintln(w, "tu as perdue 1pv il te reste", Guesses, "pv")
				Messages = "Dommage mauvaise mot"
			}
		}
	}
	axel = false
	if Show_Hide(guess, word_guess) == guess {
		// fmt.Fprintln(w, "Tu as gagné !!")
		Messages = "Tu as gagné"
		victory = 1
	}
	if GuessesLeft == Guesses2+1 {
		GuessesLeft -= 1
		//fmt.Fprintln(w, draw[10-Guesses])
		Hangman2 = Draw[10-Guesses2]

	}
	if GuessesLeft == Guesses2+2 {
		GuessesLeft -= 2
		//fmt.Fprintln(w, draw[10-Guesses])
		Hangman2 = Draw[10-Guesses2]
	}
	if Guesses2 == 0 {
		Hangman2 = Draw[10-Guesses2]
		word_guess = ""
		Guesses2 = 10
		Messages = "Tu es mort !! "
		win_next = 0
		restart()
	}
	if victory == 1 {
		win_all += 1
		win_next += 1
		restart()
		random()
		victory = 0
	}
	data := Hangman{
		Display:       Show_Hide(guess, word_guess),
		word:          "",
		Message:       Messages,
		Guesses:       Guesses2,
		state:         state,
		guess:         guess,
		draw:          Draw,
		Hangman:       Hangman2,
		WrongGuesses2: wrongGuesses,
		Win:           win_all,
		Ratio:         win_next,
	}

	tmpl.Execute(w, data)
}

func main() {
	http.HandleFunc("/jeu", css)
	http.Handle("/c", http.StripPrefix("/c", fs))
	print("Lancement de la page instancier sur : 127.0.0.1:8888")
	// premier paramètre string url, après le port
	guess = ""
	tmpl = template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", ServeFiles)
	http.HandleFunc("/game/", css)
	http.ListenAndServe(":8888", nil)
	log.Fatal(http.ListenAndServe(PORT, nil))
}
func ServeFiles(w http.ResponseWriter, r *http.Request) {
	compteur = 0
	switch r.Method {

	case "GET":

		path := r.URL.Path

		fmt.Println(path)

		if path == "/" {

			path = "index.html"
		} else {

			path = "." + path
		}

		http.ServeFile(w, r, path)

	case "POST":

		r.ParseMultipartForm(0)

		message := r.FormValue("message")

		fmt.Println("----------------------------------")
		fmt.Println("Mode de jeu choisit: ", message)
		// réponse du client
		if message == "Facile" {
			click = 0
			random()
			fmt.Fprintf(w, "tu as choisit le mode facile, pour le prochain tour")
			compteur = 1
		}
		if message == "Normale" {
			click = 2
			random()
			fmt.Fprintf(w, "tu as choisit le mode normale, pour le prochain tour")
			compteur = 1
		}
		if message == "Difficile" {
			click = 3
			random()
			fmt.Fprintf(w, "tu as choisit le mode difficile, pour le prochain tour")
			compteur = 1
		}
		if compteur == 0 {
			fmt.Fprintf(w, "Le mode choisit n'existe pas désolé")
		}
	}
}
