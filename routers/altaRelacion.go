package routers

import (
	"net/http"

	"github.com/lautarofaridtoledo/Go-microblogging-server/bd"
	"github.com/lautarofaridtoledo/Go-microblogging-server/models"
)

/*AltaRelacion - realiza el registro de la relación entre usuarios*/
func AltaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}
	var rel models.Relacion
	rel.UsuarioID = IDUsuario
	rel.UsuarioRelacionID = ID

	status, err := bd.InsertoRelacion(rel)
	//si hay un error
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar insertar relación, intentelo nuevamente. "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado insertar la relación. "+err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
