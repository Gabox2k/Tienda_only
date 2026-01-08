package handlers

import (
	"bytes"
	"context"
	"encoding/json"
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
	if len(carrito.Items) == 0 {
		http.Redirect(w, r, "/carrito", http.StatusSeeOther)
		return
	}

	var productosPayload []map[string]interface{}
	for _, item := range carrito.Items {
		productosPayload = append(productosPayload, map[string]interface{}{
			"nombre":   item.Producto.Nombre,
			"precio":   item.Producto.Precio,
			"cantidad": item.Cantidad,
		})
	}

	payload := map[string]interface{}{
		"productos": productosPayload,
		"direccion": "dirrecion por defecto",
	}

	b, err := json.Marshal(payload)
	if err != nil {
		http.Error(w, "error en el json", http.StatusInternalServerError)
		return
	}

	resp, err := http.Post("http://localhost:3000/orden/crear", "application/json", bytes.NewBuffer(b))
	if err != nil || resp.StatusCode != 200 {
		http.Error(w, "no se pudo crear el orden ", http.StatusInternalServerError)
		return
	}

	carrito.Items = []modelo.ItemCarro{}

	http.Redirect(w, r, "/ok", http.StatusSeeOther)
}
