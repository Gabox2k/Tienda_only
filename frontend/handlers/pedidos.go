package handlers

import (
	"context"
	"html/template"
	"net/http"

	"tienda_only/db"
	"tienda_only/modelo"
)

func Pedido(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/pedido.html",
	))
	tmpl.Execute(w, nil)
}

func CrearPedido(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	orden := modelo.Orden{
		Producto:  r.FormValue("producto"),
		Direccion: r.FormValue("direccion"),
	}

	db.DB.Collection("ordenes").InsertOne(context.TODO(), orden)
	http.Redirect(w, r, "/ok", http.StatusSeeOther)
}

func PedidoOK(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/ok.html",
	))

	tmpl.Execute(w, nil)
}
