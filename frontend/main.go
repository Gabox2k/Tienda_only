package main

import (
	"log"
	"net/http"
	"tienda_only/db"

	"tienda_only/rutas"
)

func main() {
	//Conexion al Mongo
	db.Conexion()

	//Accede a las imagenes
	http.Handle("/imagen/", http.StripPrefix("/imagen/", http.FileServer(http.Dir("imagen"))))

	//Carga las rutas
	rutas.CargaRutas()

	log.Println("Servidor Go en http://localhost:4000")
	http.ListenAndServe(":4000", nil)
}
