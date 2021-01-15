package main

import (
	"log"

	"github.com/lautarofaridtoledo/Go-microblogging/bd"
	handler "github.com/lautarofaridtoledo/Go-microblogging/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handler.Manejadores()
}
