package main

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"strings"
	"time"
)

type Cep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/cep/", FindCepHandler)
	mux.HandleFunc("/ping", PingHandler)

	println("Starting server at port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

func PingHandler(w http.ResponseWriter, _ *http.Request) {
	pong, err := json.Marshal(map[string]string{"ping": "pong", "status": "ok", "time": time.Now().UTC().String()})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = w.Write(pong)
	if err != nil {
		return
	}
}

func FindCepHandler(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	path := r.URL.Path
	segments := strings.Split(path, "/")
	cepCode := segments[len(segments)-1]

	if cepCode == "" {
		http.Error(w, "Cep is required", http.StatusBadRequest)
		return
	}

	cepRes, err := http.Get("https://viacep.com.br/ws/" + cepCode + "/json")
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			panic(err)
		}
	}(cepRes.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	cepInfo, err := io.ReadAll(cepRes.Body)

	var cep Cep
	err = json.Unmarshal(cepInfo, &cep)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err = t.Execute(w, cep)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
