package modelo

import "go.mongodb.org/mongo-driver/bson/primitive"

type Producto struct {
	Id     primitive.ObjectID `bson:"_id,omitempty"`
	Nombre string             `bson:"nombre"`
	Precio float64            `bson:"precio"`
}
