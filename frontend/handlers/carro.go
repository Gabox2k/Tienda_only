package handlers

import (
	"context"
	"html/template"
	"net/http"

	"tienda_only/db"
	"tienda_only/modelo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var carrito = modelo.Carrito{}

func AgregarAlCarrito(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")

	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "id no valido", http.StatusBadRequest)
		return
	}

	var producto modelo.Producto
	err = db.DB.Collection("productos").
		FindOne(context.TODO(), bson.M{"_id": objId}).
		Decode(&producto)

	if err != nil {
		http.Error(w, "producto no encontrado", http.StatusNotFound)
		return
	}

	for i, item := range carrito.Items {
		if item.Producto.Id == producto.Id {
			carrito.Items[i].Cantidad++
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	carrito.Items = append(carrito.Items, modelo.ItemCarro{
		Producto: producto,
		Cantidad: 1,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func VerCarrito(w http.ResponseWriter, r *http.Request) {

	total := 0.0
	for _, item := range carrito.Items {
		total += item.Producto.Precio * float64(item.Cantidad)
	}

	data := struct {
		Items []modelo.ItemCarro
		Total float64
	}{
		Items: carrito.Items,
		Total: total,
	}

	tmpl := template.Must(template.ParseFiles(
		"templates/layout.html",
		"templates/carrito.html",
	))

	tmpl.Execute(w, data)
}

func CancelarCarrito(w http.ResponseWriter, r *http.Request) {
	carrito.Items = []modelo.ItemCarro{}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func ComprarCarrito(w http.ResponseWriter, r *http.Request) {
	for _, item := range carrito.Items {
		orden := modelo.Orden{
			Producto:  item.Producto.Nombre,
			Direccion: "Dirrecicon por defecto",
		}

		db.DB.Collection("ordenes").InsertOne(context.TODO(), orden)

	}

	carrito.Items = []modelo.ItemCarro{}

	http.Redirect(w, r, "/ok", http.StatusSeeOther)
}
