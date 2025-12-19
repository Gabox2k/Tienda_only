package handlers

import (
	"context"
	"html/template"
	"net/http"

	"frontend/db"

	"go.mongodb.org/mongo-driver/bson"
)

func mostrarProductos(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"templats/layout.html",
		"templates/productos.html",
	))

	cursor, _ := db.DB.Collection("productos").Find(context.TODO(), bson.M{})

	var productos []modelo.producto
	cursor.All(context.TODO(), &productos)

	tmpl.Execute(w, productos)

}
