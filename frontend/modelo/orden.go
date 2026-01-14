package modelo

import "go.mongodb.org/mongo-driver/bson/primitive"

//Orden de compra
type Orden struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	Producto  string             `bson:"producto"`
	Direccion string             `bson:"direccion"`
}
