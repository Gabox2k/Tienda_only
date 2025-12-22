package rutas

import (
	"net/http"

	"tienda_only/handlers"
)

func CargaRutas() {
	http.HandleFunc("/", handlers.MostrarProductos)
	http.HandleFunc("/pedido", handlers.Pedido)
	http.HandleFunc("/pedido/crear", handlers.CrearPedido)
	http.HandleFunc("/ok", handlers.PedidoOK)

}
