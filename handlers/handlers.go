package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/lautarofaridtoledo/Go-microblogging-server/middlew"
	"github.com/lautarofaridtoledo/Go-microblogging-server/routers"
	"github.com/rs/cors"
)

//Manejadores : seteo mi puerto, el handler y pongo a escuchar al servidor*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlew.ChequeoBD(routers.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlew.ChequeoBD(middlew.ValidoJWT(routers.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.GraboTweet))).Methods("POST")
	router.HandleFunc("/leoTweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweet))).Methods("GET")
	router.HandleFunc("/borrarTweet", middlew.ChequeoBD(middlew.ValidoJWT(routers.EliminarTweet))).Methods("DELETE")
	//-Imagenes
	router.HandleFunc("/subirAvatar", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlew.ChequeoBD(routers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/subirBanner", middlew.ChequeoBD(middlew.ValidoJWT(routers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerBanner", middlew.ChequeoBD(routers.ObtenerBanner)).Methods("GET")

	//-Relaciones
	router.HandleFunc("/altaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.AltaRelacion))).Methods("POST")
	router.HandleFunc("/bajaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.BajaRelacion))).Methods("DELETE")
	router.HandleFunc("/consultaRelacion", middlew.ChequeoBD(middlew.ValidoJWT(routers.ConsultaRelacion))).Methods("GET")

	//-Lista Usuarios
	router.HandleFunc("/listaUsuarios", middlew.ChequeoBD(middlew.ValidoJWT(routers.ListaUsuarios))).Methods("GET")

	//-Mostrar Tweets Seguidores
	router.HandleFunc("/leoTweetsSeguidores", middlew.ChequeoBD(middlew.ValidoJWT(routers.LeoTweetsSeguidores))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
