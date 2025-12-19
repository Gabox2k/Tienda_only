package modelo

import "go.mongodb.org/mongo-driver/bson/primitive"

type prodcuto struct {
	id     primitive.ObjectID `bson:"_id,omitempty"`
	nombre string             `bson:"nombre"`
	precio float64            `bson:"precio"`
}
