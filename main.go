package main

import (
	"log"

	"github.com/lautarofaridtoledo/Go-microblogging-server/bd"
	handler "github.com/lautarofaridtoledo/Go-microblogging-server/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handler.Manejadores()
}
