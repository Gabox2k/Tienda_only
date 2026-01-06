package modelo

type ItemCarro struct {
	Producto Producto
	Cantidad int
}

type Carrito struct {
	Items []ItemCarro
}
