package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//Color is a main color schem
type Color struct {
	Name          string `json:"name"`
	Transcription string `json:"transcription"`
	Description   string `json:"description"`
	Hex           string `json:"hex"`
}

//GetColors for export colors
func GetColors() []Color {
	var Colors = []Color{
		Color{"black", "БЛЭК", "чёрный", "#000"},
		Color{"red", "РЕД", "красный", "#f00"},
		Color{"yellow", "ЕЛОУ", "жёлтый", "#ff0"},
		Color{"blue", "БЛУ", "синий", "#00f"},
		Color{"green", "ГРИН", "зелёный", "#0f0"},
	}
	return Colors
}

func main() {
	port := ":8080"

	http.Handle("/", http.FileServer(http.Dir("./src")))
	http.HandleFunc("/api/colors", getColors)

	log.Println("Server working on " + port + " port")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("server error: ", err)
		return
	}
}

func getColors(w http.ResponseWriter, r *http.Request) {
	colors := GetColors()
	jData, err := json.Marshal(colors)

	if err != nil {
		fmt.Println("Ошибка")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)
}
