package routers

import (
	"net/http"

	"github.com/lautarofaridtoledo/Go-microblogging-server/bd"
)

/*EliminarTweet - elimina un tweet del usuario*/
func EliminarTweet(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}
	err := bd.BorroTweet(ID, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar el tweet"+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("context-type", "encoding/json")
	w.WriteHeader(http.StatusCreated)
}
