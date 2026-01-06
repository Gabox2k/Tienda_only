package main

import (
	"log"
	"net/http"
	"tienda_only/db"

	"tienda_only/rutas"
)

func main() {
	db.Conexion()

	http.Handle("/imagen/", http.StripPrefix("/imagen/", http.FileServer(http.Dir("frontend/imagen"))))

	rutas.CargaRutas()

	log.Println("Servidor Go en http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}
