package handlers

import (
	"context"
	"fmt"
	"html/template"
	"net/http"

	"tienda_only/db"
	"tienda_only/modelo"

	"go.mongodb.org/mongo-driver/bson"
)

func precio(valor float64) string {
	s := fmt.Sprintf("%.0f", valor)

	n := len(s)
	if n <= 3 {
		return s
	}

	return s[:n-3] + "." + s[n-3:]
}

func MostrarProductos(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("layout.html").
		Funcs(template.FuncMap{
			"precio": precio,
		}).ParseFiles(
		"templates/layout.html",
		"templates/productos.html"),
	)

	cursor, _ := db.DB.Collection("productos").Find(context.TODO(), bson.M{})

	var productos []modelo.Producto
	cursor.All(context.TODO(), &productos)

	tmpl.Execute(w, productos)
}
