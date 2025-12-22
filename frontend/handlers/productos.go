package handlers

import (
	"context"
	"html/template"
	"net/http"

	"tienda_only/db"
	"tienda_only/modelo"

	"go.mongodb.org/mongo-driver/bson"
)

func MostrarProductos(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/productos.html",
	))

	cursor, _ := db.DB.Collection("productos").Find(context.TODO(), bson.M{})

	var productos []modelo.Producto
	cursor.All(context.TODO(), &productos)

	tmpl.Execute(w, productos)
}
