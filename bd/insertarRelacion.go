package bd

import (
	"context"
	"time"

	"github.com/lautarofaridtoledo/Go-microblogging-server/models"
)

/*InsertoRelacion - graba la relación en la bd*/
func InsertoRelacion(relac models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("Microblogging-Go")
	col := db.Collection("relacion")

	//Directamente trato de insertar en la base de datos el modelo que recibo
	_, err := col.InsertOne(ctx, relac)
	if err != nil {
		return false, err
	}
	return true, nil
}
