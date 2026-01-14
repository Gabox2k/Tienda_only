package handlers

import (
	"context"
	"html/template"
	"net/http"

	"tienda_only/db"
	"tienda_only/modelo"

	"go.mongodb.org/mongo-driver/bson"
)

// Muestra el pedido
func Pedido(w http.ResponseWriter, r *http.Request) {
	//Obtiene el id
	id := r.URL.Query().Get("id")

	var producto modelo.Producto

	//Busca el producto por id
	db.DB.Collection("productos").
		FindOne(context.TODO(), bson.M{"_id": id}).
		Decode(&producto)

	tmpl := template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/pedido.html",
	))

	tmpl.Execute(w, producto)
}

// Crea el pedido
func CrearPedido(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	id := r.FormValue("id")

	var producto modelo.Producto

	//Busca el producto por id
	db.DB.Collection("productos").
		FindOne(context.TODO(), bson.M{"_id": id}).
		Decode(&producto)

	//Crea un orden
	orden := modelo.Orden{
		Producto:  r.FormValue("producto"),
		Direccion: r.FormValue("direccion"),
	}

	//Insera el orden en ordenes
	db.DB.Collection("ordenes").InsertOne(context.TODO(), orden)
	http.Redirect(w, r, "/ok", http.StatusSeeOther)
}

// Confirmacion
func PedidoOK(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/ok.html",
	))

	tmpl.Execute(w, nil)
}
