package modelo

import "go.mongodb.org/mongo-driver/bson/primitive"

type orden struct {
	id        primitive.ObjectID `bson:"_id,omitempty"`
	prodcuto  string             `bson:"producto"`
	direccion string             `bson:"direccion"`
}
