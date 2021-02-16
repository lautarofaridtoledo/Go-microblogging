package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/lautarofaridtoledo/Go-microblogging-server/bd"
	"github.com/lautarofaridtoledo/Go-microblogging-server/models"
)

/*SubirBanner sube el Banner al servidor */
func SubirBanner(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("banner")

	// Se extrae la extensión del archivo original del banner
	var extension = strings.Split(handler.Filename, ".")[1]

	/*En carpeta la banners, se encontrara la imagen renombrada con como nombre IDUsuario,
	La carpeta banners fue previamente creada.*/
	var archivo string = "uploads/banners/" + IDUsuario + "." + extension

	//Creo el manejador de archivo (lectura, escritura y ejecución).
	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Error al subir la imagen "+err.Error(), http.StatusBadRequest)
		return
	}
	//Copio la imagen en f (file)
	_, err = io.Copy(f, file)
	if err != nil {
		http.Error(w, "Error al copiar la imagen "+err.Error(), http.StatusBadRequest)
		return
	}

	var usuario models.Usuario
	var status bool
	//Renombro el banner.
	usuario.Banner = IDUsuario + "." + extension

	//Grabo en la bd el cambio en el campo banner.
	status, err = bd.ModificoRegistro(usuario, IDUsuario)
	if err != nil || status == false {
		http.Error(w, "Error al grabar el banner en la bd "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	//doy  status created.
	w.WriteHeader(http.StatusCreated)
}
