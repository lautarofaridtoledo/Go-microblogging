package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/lautarofaridtoledo/Go-microblogging/bd"
)

//LeoTweet se encarga de leer los tweets almacenados en la BD
func LeoTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	// si todo anduvo bien:
	if len(r.URL.Query().Get("pagina")) < 1 {
		http.Error(w, "debe enviar el parámetro pagina", http.StatusBadRequest)
		return
	}

	// Vamos a trabajar con la paginación
	pag, err := strconv.Atoi(r.URL.Query().Get("pagina"))
	if err != nil {
		http.Error(w, "debe enviar el parámetro pagina con un valor mayor a cero", http.StatusBadRequest)
		return
	}

	pagina := int64(pag)

	tweets, result := bd.LeoTweets(ID, pagina)
	if result == false {
		http.Error(w, "error al leer los Tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("context-type", "encoding/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tweets)
}
