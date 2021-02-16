package routers

import (
	"net/http"

	"github.com/lautarofaridtoledo/Go-microblogging-server/bd"
	"github.com/lautarofaridtoledo/Go-microblogging-server/models"
)

/*BajaRelacion - realiza el borrado de la relación entre usuarios*/
func BajaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	var rel models.Relacion
	rel.UsuarioID = IDUsuario
	rel.UsuarioRelacionID = ID

	status, err := bd.BorroRelacion(rel)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar borrar relacion "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "No se ha logrado borrar la relación. "+err.Error(), http.StatusBadRequest)
		return
	}
}
