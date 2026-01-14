import mongoose from "mongoose"

//Guarda el orden en Mongo
const productoScheam = new mongoose.Schema({
    nombre: String,
    precio: Number,
    stock: Number
})

export default mongoose.model("Producto", productoScheam)