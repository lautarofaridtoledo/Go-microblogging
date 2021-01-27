package bd

import (
	"context"
	"time"

	"github.com/lautarofaridtoledo/Go-microblogging/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//LeoTweets recibe el ID del tweet
func LeoTweets(id string, pagina int64) ([]*models.DevuelvoTweets, bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("Microblogging-Go")
	col := db.Collection("tweet")

	var tweets []*models.DevuelvoTweets

	//La condicion es el ID del tweet (Cada tweet en la MongoDB tiene su propio ID con un ID del usuario que lo escribió)
	condicion := bson.M{
		"userid": id,
	}

	opciones := options.Find()

	opciones.SetLimit(20)

	//Traerá los tweets ordenados de forma descendente (indicado por el valor de value -1)
	opciones.SetSort(bson.D{{Key: "fecha", Value: -1}})

	//indico de a cuantos documentos tiene que ir salteando.
	opciones.SetSkip((pagina - 1) * 20)

	cursor, err := col.Find(ctx, condicion, opciones)

	if err != nil {
		return tweets, false
	}

	for cursor.Next(context.TODO()) {

		var registro models.DevuelvoTweets

		er := cursor.Decode(&registro)

		if er != nil {
			return tweets, false
		}

		tweets = append(tweets, &registro)
	}

	return tweets, true
}
