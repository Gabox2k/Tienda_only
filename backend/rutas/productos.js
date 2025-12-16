import mongoose from "mongoose"

const productoScheam = new mongoose.Schema({
    nombre: String,
    precio: Number,
    stock: Number
})

export default mongoose.model("Producto", productoScheam)