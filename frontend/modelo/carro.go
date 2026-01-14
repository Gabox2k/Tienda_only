package modelo

//Define el item en el carro
type ItemCarro struct {
	Producto Producto
	Cantidad int
}

//Carrito completo
type Carrito struct {
	Items []ItemCarro
}
