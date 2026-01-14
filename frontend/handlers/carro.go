package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"

	"tienda_only/db"
	"tienda_only/modelo"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var carrito = modelo.Carrito{}

// Funcion para agregar los productos al carrito
func AgregarAlCarrito(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	id := r.FormValue("id")

	//Covierte el id a ObjId
	objId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "id no valido", http.StatusBadRequest)
		return
	}

	var producto modelo.Producto

	//Busca el producto por su id
	err = db.DB.Collection("productos").
		FindOne(context.TODO(), bson.M{"_id": objId}).
		Decode(&producto)

	//Error cuando no se encuentra el producto
	if err != nil {
		http.Error(w, "producto no encontrado", http.StatusNotFound)
		return
	}

	//Recorre los productos del carrito
	for i, item := range carrito.Items {
		if item.Producto.Id == producto.Id {
			carrito.Items[i].Cantidad++
			http.Redirect(w, r, "/", http.StatusSeeOther)
			return
		}
	}

	//Agrega un producto al carrito
	carrito.Items = append(carrito.Items, modelo.ItemCarro{
		Producto: producto,
		Cantidad: 1,
	})

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Formatea el numero
func precio2(valor float64) string {
	s := fmt.Sprintf("%.0f", valor)

	n := len(s)
	if n <= 3 {
		return s
	}

	return s[:n-3] + "." + s[n-3:]
}

// Muestra lo que hay en el carrito
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

	tmpl := template.Must(template.New("layout.html").
		Funcs(template.FuncMap{
			"precio": precio2,
		}).
		ParseFiles(
			"templates/layout.html",
			"templates/carrito.html"),
	)
	tmpl.Execute(w, data)
}

// Cancelar la compra
func CancelarCarrito(w http.ResponseWriter, r *http.Request) {
	carrito.Items = []modelo.ItemCarro{}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Compra la que hay dentro del carrito
func ComprarCarrito(w http.ResponseWriter, r *http.Request) {
	if len(carrito.Items) == 0 {
		http.Redirect(w, r, "/carrito", http.StatusSeeOther)
		return
	}

	//Convierte en formato JSON
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

	//Envia al backend
	resp, err := http.Post("http://localhost:3000/orden/crear", "application/json", bytes.NewBuffer(b))
	if err != nil || resp.StatusCode != 200 {
		http.Error(w, "no se pudo crear el orden ", http.StatusInternalServerError)
		return
	}

	carrito.Items = []modelo.ItemCarro{}

	http.Redirect(w, r, "/ok", http.StatusSeeOther)
}
