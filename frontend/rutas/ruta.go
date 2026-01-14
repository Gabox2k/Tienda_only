package rutas

import (
	"net/http"

	"tienda_only/handlers"
)

// Registra las rutas
func CargaRutas() {
	http.HandleFunc("/", handlers.MostrarProductos)

	http.HandleFunc("/carrito", handlers.VerCarrito)
	http.HandleFunc("/carrito/agregar", handlers.AgregarAlCarrito)
	http.HandleFunc("/carrito/cancelar", handlers.CancelarCarrito)
	http.HandleFunc("/carrito/comprar", handlers.ComprarCarrito)

	http.HandleFunc("/pedido", handlers.Pedido)
	http.HandleFunc("/pedido/crear", handlers.CrearPedido)
	http.HandleFunc("/ok", handlers.PedidoOK)

}
